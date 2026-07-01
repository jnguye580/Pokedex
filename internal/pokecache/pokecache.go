package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	sync  sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{cache: make(map[string]cacheEntry)}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for range ticker.C {
		c.sync.Lock()
		for key, val := range c.cache {
			if time.Since(val.createdAt) > interval {
				delete(c.cache, key)
			}
		}
		c.sync.Unlock()
	}

}

func (c *Cache) Add(key string, val []byte) {
	c.sync.Lock()
	defer c.sync.Unlock()

	cacheE := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache[key] = cacheE
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.sync.Lock()
	defer c.sync.Unlock()

	cacheEn, ok := c.cache[key]
	if ok {
		return cacheEn.val, true
	}

	return nil, false
}
