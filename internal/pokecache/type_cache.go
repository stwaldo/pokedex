package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	lock sync.Mutex
	duration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	
	entry, exists := c.entries[key]
	if !exists {
		return nil, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop() {
	interval := c.duration
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	
	for {
		<-ticker.C
		c.lock.Lock()
		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.duration {
				delete(c.entries, key)
			}
		}
		c.lock.Unlock()
	}

} 

func NewCache(duration time.Duration) *Cache {
	c := &Cache{
		entries: make(map[string]cacheEntry),
		duration: duration,
	}
	go c.reapLoop()
	return c
}