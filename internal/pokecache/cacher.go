package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Cached map[string]CacheEntry
	Mu     sync.Mutex
}

type CacheEntry struct {
	CreatedAt    time.Time
	Data         []byte
	NextPage     string
	PreviousPage string
}

func NewCache() *Cache {
	cache := &Cache{
		Cached: make(map[string]CacheEntry),
	}
	return cache
}

func (c *Cache) Add(key, nextPage, previousPage string, data []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	createdAt := time.Now()
	c.Cached[key] = CacheEntry{
		CreatedAt:    createdAt,
		Data:         data,
		NextPage:     nextPage,
		PreviousPage: previousPage,
	}
}

func (c *Cache) Get(key string) (CacheEntry, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	var cached CacheEntry
	var blankCache CacheEntry
	cached = c.Cached[key]
	if cached.Data == nil {
		return blankCache, false
	}
	fmt.Printf("\nGet() returns data: %v", cached.Data)
	return cached, true
}
