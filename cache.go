package cache

import (
	"time"
)

const (
	// NoExpiration Item has no expiration date
	NoExpiration time.Duration = -1
	// DefaultExpiration Equivalent to passing in the same expiration duration as was given to New() or NewFrom()
	// when the cache was created (e.g. 5 minutes.)
	DefaultExpiration time.Duration = 0
)

func New[K comparable, V any](defaultExpiration, cleanupInterval time.Duration) *Any[K, V] {
	items := make(map[K]Item[V])
	return newCacheAnyWithJanitor(defaultExpiration, cleanupInterval, items)
}

func NewFrom[K comparable, V any](defaultExpiration, cleanupInterval time.Duration, items map[K]Item[V]) *Any[K, V] {
	return newCacheAnyWithJanitor(defaultExpiration, cleanupInterval, items)
}

func NewAny[K comparable, V any](defaultExpiration, cleanupInterval time.Duration) *Any[K, V] {
	items := make(map[K]Item[V])
	return newCacheAnyWithJanitor(defaultExpiration, cleanupInterval, items)
}

func NewAnyFrom[K comparable, V any](defaultExpiration, cleanupInterval time.Duration, items map[K]Item[V]) *Any[K, V] {
	return newCacheAnyWithJanitor(defaultExpiration, cleanupInterval, items)
}

func NewNumber[K comparable, V number](defaultExpiration, cleanupInterval time.Duration) *Number[K, V] {
	items := make(map[K]Item[V])
	return newCacheNumberWithJanitor(defaultExpiration, cleanupInterval, items)
}

func NewNumberFrom[K comparable, V number](defaultExpiration, cleanupInterval time.Duration, items map[K]Item[V]) *Number[K, V] {
	return newCacheNumberWithJanitor(defaultExpiration, cleanupInterval, items)
}
