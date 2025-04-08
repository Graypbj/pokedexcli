package pokecache

import (
	"sync"
	"time"
)

// Cache
type Cache struct {
	mu           sync.Mutex
	cacheEntries map[string]cacheEntry
	interval     time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheEntries: make(map[string]cacheEntry),
		interval:     interval,
	}

	go cache.reapLoop()

	return cache
}

// Add
func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	return
}

// Get
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.cacheEntries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for key, item := range c.cacheEntries {
			if time.Since(item.createdAt) > c.interval {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}
}
