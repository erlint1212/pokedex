package pokecache

import (
    "time"
    "sync"
)

type Cache struct {
    mu          sync.Mutex
    v           map[string]cacheEntry
    interval    time.Duration
}

type cacheEntry struct {
    createdAt   time.Time
    val         []byte
}

func NewCache(interval time.Duration) Cache {
    cache := Cache{
        v: map[string]cacheEntry{},
    }
    cache.reapLoop(interval)
    return cache
}

func (c *Cache) Add(key string, val []byte) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.v[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.mu.Lock()
    defer c.mu.Unlock()
    v, OK := c.v[key]
    if !OK {
        return []byte{}, false
    }
    return v.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
    go func() {
        for {
            current_time := time.Now()
            time.Sleep(interval)
            c.mu.Lock()
            for key, cache_entry := range c.v {
                if cache_entry.createdAt.Before(current_time.Add(interval)) {
                    delete(c.v, key)
                }
            }
            c.mu.Unlock()
        }
    }()
}

