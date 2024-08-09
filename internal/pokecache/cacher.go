package pokecache

import (
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
