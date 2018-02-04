package main

import "sync"

type Data interface{}

type Cache interface {
	Load(key string) (*Data, bool)
	Store(key string, d *Data)
}

type SimpleCache struct {
	lock sync.RWMutex
	data map[string]Data
}

func (c *SimpleCache) Load(key string) (*Data, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	d, ok := c.data[key]
	return &d, ok
}

func (c *SimpleCache) Store(key string, d *Data) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = *d
}
