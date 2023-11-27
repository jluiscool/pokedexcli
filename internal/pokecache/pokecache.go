package pokeacache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

// a cache entry is an http body, which is stored as bytes
// created at allows us to purge catch in case we want to refetch the data
type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

// helper method to create new cache
// kind of like a constructor
func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

// adds new entries to cache, like a constructor
// note that it uses a pointer to it directly modifies the Cache receiver
func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

// get function
// bool tells whether that key exists or not
func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	//ticker.C is the channel the ticker operates on
	//if ticker is 5 minutes, for loop runs every 5 mins
	for range ticker.C {
		c.reap(interval)
	}
}

// purge cache after a certain interval
func (c *Cache) reap(interval time.Duration) {
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
