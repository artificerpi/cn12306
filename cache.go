package main

import "sync"

// Data is the value type of cache.
type Data interface{}

// The Cache interface represents an entity you can load and store data.
type Cache interface {
	Load(key string) (*Data, bool)
	Store(key string, d *Data)
}

// SimpleCache is a simple implementation of Cache with a read-write lock.SimpleCache
// Note that rw-lock is locked on the methods rather than bucket of map
type SimpleCache struct {
	sync.RWMutex
	data map[string]Data
}

// Load loads data from cache by key
func (c *SimpleCache) Load(key string) (*Data, bool) {
	c.RLock()
	defer c.RUnlock()
	d, ok := c.data[key]
	return &d, ok
}

// Store saves data to cache with specified key
func (c *SimpleCache) Store(key string, d *Data) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = *d
}
