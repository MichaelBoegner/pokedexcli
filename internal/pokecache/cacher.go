package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	Cached map[string]cacheEntry
	Mu     sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func NewCache() *Cache {
	cache := &Cache{
		Cached: make(map[string]cacheEntry),
	}
	return cache
}

func (c *Cache) Add(key string, data []byte) {

	createdAt := time.Now()
	c.Cached[key] = cacheEntry{
		createdAt: createdAt,
		data:      data,
	}
	fmt.Printf("\ncache of c: %v", c)
}

func (c *Cache) Get(key string) ([]byte, bool) {
	var data []byte
	data = c.Cached[key].data
	if data == nil {
		return nil, false
	}
	return data, true
}
