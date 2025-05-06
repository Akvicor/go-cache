package cache

import (
	"runtime"
	"time"
)

// newCacheAnyWithJanitor create new cache with janitor
func newCacheAnyWithJanitor[K comparable, V any](de time.Duration, ci time.Duration, m map[K]Item[V]) *Any[K, V] {
	c := newCache(de, m)
	C := &Any[K, V]{c}
	if ci > 0 {
		runJanitor(c, ci)
		runtime.SetFinalizer(C, stopJanitor)
	}
	return C
}

type Any[K comparable, V any] struct {
	*cache[K, V]
}

