package cache

import "sync"

type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func New() *Cache {
	return &Cache{store: make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	c.store[key] = value
	c.mu.Unlock()
}

func (c *Cache) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.store[key]
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	delete(c.store, key)
	c.mu.Unlock()
}
