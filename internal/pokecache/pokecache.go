package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mu    sync.Mutex
	Cache map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.Cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	defer c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.Cache[key]
	defer c.mu.Unlock()
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.Reap(time.Now(), interval)
	}
}

func (c *Cache) Reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	for key := range c.Cache {
		if time.Since(c.Cache[key].createdAt) > interval {
			delete(c.Cache, key)
		}
	}
	defer c.mu.Unlock()
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{Cache: map[string]cacheEntry{}}
	go cache.reapLoop(interval)
	return &cache
}
