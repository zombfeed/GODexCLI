package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) Reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	for key, value := range c.Cache {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.Cache, key)
		}
	}
	defer c.mu.Unlock()
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		Cache: map[string]cacheEntry{},
		mu:    &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}
