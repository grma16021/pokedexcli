package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

//now := time.Now()

func NewCache(interval time.Duration) *Cache {

	c := &Cache{
		entries: make(map[string]cacheEntry),
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cach := cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}

	c.entries[key] = cach

}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	} else {
		return nil, false
	}

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.mu.Lock()
		for i, value := range c.entries {
			elapsed := time.Since(value.createdAt)
			if elapsed > interval {

				delete(c.entries, i)
			}
		}
		c.mu.Unlock()

	}

}
