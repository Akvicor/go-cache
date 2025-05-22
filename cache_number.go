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

// Increment an item by n. Returns an error if the
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

// Decrement an item by n. Returns an error if the
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

// SetMax Update Value to the maximum value. Create it if it does not exist. If the duration is 0
// (DefaultExpiration), the cache's default expiration time is used. If it is -1
// (NoExpiration), the item never expires.
func (c *Number[K, V]) SetMax(k K, v V, d time.Duration) error {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	item, found := c.items[k]
	if !found || item.Expired() {
		c.items[k] = Item[V]{
			Value:      v,
			Expiration: e,
			Hit:        0,
		}
		return nil
	}
	item.Value = max(item.Value, v)
	c.items[k] = item
	return nil
}

// SetMin Update Value to the minimum value. Create it if it does not exist. If the duration is 0
// (DefaultExpiration), the cache's default expiration time is used. If it is -1
// (NoExpiration), the item never expires.
func (c *Number[K, V]) SetMin(k K, v V, d time.Duration) error {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	item, found := c.items[k]
	if !found || item.Expired() {
		c.items[k] = Item[V]{
			Value:      v,
			Expiration: e,
			Hit:        0,
		}
		return nil
	}
	item.Value = min(item.Value, v)
	c.items[k] = item
	return nil
}

// UpdateMax Update Value to the maximum value.
func (c *Number[K, V]) UpdateMax(k K, v V) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, found := c.items[k]
	if !found || item.Expired() {
		return fmt.Errorf("Item %v not found", k)
	}
	item.Value = max(item.Value, v)
	c.items[k] = item
	return nil
}

// UpdateMin Update Value to the minimum value.
func (c *Number[K, V]) UpdateMin(k K, v V) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	item, found := c.items[k]
	if !found || item.Expired() {
		return fmt.Errorf("Item %v not found", k)
	}
	item.Value = min(item.Value, v)
	c.items[k] = item
	return nil
}
