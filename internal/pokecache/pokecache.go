package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mux     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (c Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (c Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	for range ticker.C {
		for key, entry := range c.entries {
			if entry.createdAt.Add(duration).Before(time.Now()) {
				c.mux.Lock()
				delete(c.entries, key)
				c.mux.Unlock()
			}
		}
	}
}

func NewCache(duration time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mux:     &sync.Mutex{},
	}
	go c.reapLoop(duration)

	return c
}
