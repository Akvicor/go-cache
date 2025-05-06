package cache

import (
	"fmt"
	"runtime"
	"time"
)

// newCacheNumberWithJanitor create new cache with janitor
func newCacheNumberWithJanitor[K comparable, V number](de time.Duration, ci time.Duration, m map[K]Item[V]) *Number[K, V] {
	c := newCache(de, m)
	C := &Number[K, V]{c}
	if ci > 0 {
		runJanitor(c, ci)
		runtime.SetFinalizer(C, stopJanitor)
	}
	return C
}

type Number[K comparable, V number] struct {
	*cache[K, V]
}

// Increment an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to increment it by n. To retrieve the incremented value, use one
// of the specialized methods, e.g. IncrementInt64.
func (c *Number[K, V]) Increment(k K, n V) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("Item %v not found", k)
	}
	v.Value = v.Value + n
	c.items[k] = v
	c.mu.Unlock()
	return nil
}

// Decrement an item of type int, int8, int16, int32, int64, uintptr, uint,
// uint8, uint32, or uint64, float32 or float64 by n. Returns an error if the
// item's value is not an integer, if it was not found, or if it is not
// possible to decrement it by n. To retrieve the decremented value, use one
// of the specialized methods, e.g. DecrementInt64.
func (c *Number[K, V]) Decrement(k K, n V) error {
	c.mu.Lock()
	v, found := c.items[k]
	if !found || v.Expired() {
		c.mu.Unlock()
		return fmt.Errorf("Item %v not found", k)
	}
	v.Value = v.Value - n
	c.items[k] = v
	c.mu.Unlock()
	return nil
}
