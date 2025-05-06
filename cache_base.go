package cache

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// newCache create new Cache
func newCache[K comparable, V any](d time.Duration, m map[K]Item[V]) *cache[K, V] {
	if d == 0 {
		d = NoExpiration
	}
	c := &cache[K, V]{
		defaultExpiration: d,
		items:             m,
	}
	return c
}

type cache[K comparable, V any] struct {
	items             map[K]Item[V]
	mu                sync.RWMutex
	onEvicted         func(key K, value V, hit int)
	defaultExpiration time.Duration
	janitor           *janitor // Auto Clean expired item
}

func (c *cache[K, V]) OnEvicted(f func(key K, value V, hit int)) {
	c.mu.Lock()
	c.onEvicted = f
	c.mu.Unlock()
}

func (c *cache[K, V]) SetJanitor(j *janitor) {
	c.janitor = j
}

func (c *cache[K, V]) StopJanitor() {
	c.janitor.stop <- true
}

// Set Add an item to the cache, replacing any existing item. If the duration is 0
// (DefaultExpiration), the cache's default expiration time is used. If it is -1
// (NoExpiration), the item never expires.
func (c *cache[K, V]) Set(k K, v V, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[k] = Item[V]{
		Value:      v,
		Expiration: e,
		Hit:        0,
	}
}

// SetDefault Add an item to the cache, replacing any existing item, using the default expiration.
func (c *cache[K, V]) SetDefault(k K, v V) {
	c.Set(k, v, DefaultExpiration)
}

func (c *cache[K, V]) set(k K, v V, d time.Duration) {
	var e int64
	if d == DefaultExpiration {
		d = c.defaultExpiration
	}
	if d > 0 {
		e = time.Now().Add(d).UnixNano()
	}
	c.items[k] = Item[V]{
		Value:      v,
		Expiration: e,
		Hit:        0,
	}
}

func (c *cache[K, V]) get(k K) (V, bool) {
	var v V
	item, found := c.items[k]
	if !found {
		return v, false
	}
	// "Inlining" of Expired
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return v, false
		}
	}
	return item.Value, true
}

// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (c *cache[K, V]) Add(k K, v V, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if found {
		c.mu.Unlock()
		return fmt.Errorf("Item %v already exists", k)
	}
	c.set(k, v, d)
	c.mu.Unlock()
	return nil
}

// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (c *cache[K, V]) Replace(k K, x V, d time.Duration) error {
	c.mu.Lock()
	_, found := c.get(k)
	if !found {
		c.mu.Unlock()
		return fmt.Errorf("Item %v doesn't exist", k)
	}
	c.set(k, x, d)
	c.mu.Unlock()
	return nil
}

// Get an item from the cache. Returns the item or nil, and a bool indicating
// whether the key was found.
func (c *cache[K, V]) Get(k K) (V, bool) {
	var v V
	c.mu.RLock()
	// "Inlining" of get and Expired
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return v, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return v, false
		}
	}
	c.mu.RUnlock()
	return item.Value, true
}

// GetWithExpiration returns an item and its expiration time from the cache.
// It returns the item or nil, the expiration time if one is set (if the item
// never expires a zero value for time.Time is returned), and a bool indicating
// whether the key was found.
func (c *cache[K, V]) GetWithExpiration(k K) (V, time.Time, bool) {
	var v V
	c.mu.RLock()
	// "Inlining" of get and Expired
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return v, time.Time{}, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return v, time.Time{}, false
		}

		// Return the item and the expiration time
		c.mu.RUnlock()
		return item.Value, time.Unix(0, item.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	c.mu.RUnlock()
	return item.Value, time.Time{}, true
}

func (c *cache[K, V]) GetWithHit(k K) (V, int, bool) {
	var v V
	c.mu.RLock()
	// "Inlining" of get and Expired
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return v, 0, false
	}

	c.mu.RUnlock()
	return item.Value, item.Hit, true
}

func (c *cache[K, V]) GetWithHitExpiration(k K) (V, int, time.Time, bool) {
	var v V
	c.mu.RLock()
	// "Inlining" of get and Expired
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		return v, 0, time.Time{}, false
	}

	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			c.mu.RUnlock()
			return v, 0, time.Time{}, false
		}

		// Return the item and the expiration time
		c.mu.RUnlock()
		return item.Value, item.Hit, time.Unix(0, item.Expiration), true
	}

	// If expiration <= 0 (i.e. no expiration time set) then return the item
	// and a zeroed time.Time
	c.mu.RUnlock()
	return item.Value, item.Hit, time.Time{}, true
}

// DeleteExpired delete all expired items
func (c *cache[K, V]) DeleteExpired() {
	var evictedItems []keyAndValueModel[K, V]
	now := time.Now().UnixNano()
	c.mu.Lock()
	for k, v := range c.items {
		// "Inlining" of expired
		if v.Expiration > 0 && now > v.Expiration {
			ov, oh, evicted := c.delete(k)
			if evicted {
				evictedItems = append(evictedItems, keyAndValueModel[K, V]{k, ov, oh})
			}
		}
	}
	c.mu.Unlock()
	for _, v := range evictedItems {
		c.onEvicted(v.key, v.value, v.hit)
	}
}

func (c *cache[K, V]) delete(k K) (V, int, bool) {
	if c.onEvicted != nil {
		if v, found := c.items[k]; found {
			delete(c.items, k)
			return v.Value, v.Hit, true
		}
	}
	delete(c.items, k)
	var v V
	return v, 0, false
}

// Delete an item from the cache. Does nothing if the key is not in the cache.
func (c *cache[K, V]) Delete(k K) {
	c.mu.Lock()
	v, hit, evicted := c.delete(k)
	c.mu.Unlock()
	if evicted {
		c.onEvicted(k, v, hit)
	}
}

// Save Write the cache's items (using Gob) to an io.Writer.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache[K, V]) Save(w io.Writer) (err error) {
	enc := gob.NewEncoder(w)
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("Error registering item types with Gob library")
		}
	}()
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, v := range c.items {
		gob.Register(v.Value)
	}
	err = enc.Encode(&c.items)
	return
}

// Save the cache's items to the given filename, creating the file if it
// doesn't exist, and overwriting it if it does.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache[K, V]) SaveFile(fname string) error {
	fp, err := os.Create(fname)
	if err != nil {
		return err
	}
	err = c.Save(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}

// Add (Gob-serialized) cache items from an io.Reader, excluding any items with
// keys that already exist (and haven't expired) in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache[K, V]) Load(r io.Reader) error {
	dec := gob.NewDecoder(r)
	items := map[K]Item[V]{}
	err := dec.Decode(&items)
	if err == nil {
		c.mu.Lock()
		defer c.mu.Unlock()
		for k, v := range items {
			ov, found := c.items[k]
			if !found || ov.Expired() {
				c.items[k] = v
			}
		}
	}
	return err
}

// Load and add cache items from the given filename, excluding any items with
// keys that already exist in the current cache.
//
// NOTE: This method is deprecated in favor of c.Items() and NewFrom() (see the
// documentation for NewFrom().)
func (c *cache[K, V]) LoadFile(fname string) error {
	fp, err := os.Open(fname)
	if err != nil {
		return err
	}
	err = c.Load(fp)
	if err != nil {
		fp.Close()
		return err
	}
	return fp.Close()
}

// Copies all unexpired items in the cache into a new map and returns it.
func (c *cache[K, V]) Items() map[K]Item[V] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[K]Item[V], len(c.items))
	now := time.Now().UnixNano()
	for k, v := range c.items {
		// "Inlining" of Expired
		if v.Expiration > 0 {
			if now > v.Expiration {
				continue
			}
		}
		m[k] = v
	}
	return m
}

// Returns the number of items in the cache. This may include items that have
// expired, but have not yet been cleaned up.
func (c *cache[K, V]) ItemCount() int {
	c.mu.RLock()
	n := len(c.items)
	c.mu.RUnlock()
	return n
}

// Delete all items from the cache.
func (c *cache[K, V]) Flush() {
	c.mu.Lock()
	c.items = map[K]Item[V]{}
	c.mu.Unlock()
}
