package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	mu    sync.Mutex
	Cache map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Cache[key]
	if !ok {
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for {
		for key := range c.Cache {
			if time.Since(c.Cache[key].createdAt) > interval {
				delete(c.Cache, key)
			}
			fmt.Printf("%s is safe\n", key)
		}
	}
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{Cache: map[string]cacheEntry{}}
	go cache.reapLoop(interval)
	return &cache
}
