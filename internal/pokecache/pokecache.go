package pokeacache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

// a cache entry is an http body, which is stored as bytes
// created at allows us to purge catch in case we want to refetch the data
type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

// helper method to create new cache
// kind of like a constructor
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	// reapLoop needs to run in a go routine, else it runs forever, making the tests get stuck
	//can't run on the main thread
	go c.reapLoop(interval)
	return c
}

// adds new entries to cache, like a constructor
// note that it uses a pointer to it directly modifies the Cache receiver
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now().UTC(),
	}
}

// get function
// bool tells whether that key exists or not
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}

// dont block the loop
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	//ticker.C is the channel the ticker operates on
	//if ticker is 5 minutes, for loop runs every 5 mins
	for range ticker.C {
		c.reap(interval)
	}
}

// block the function that the loop runs
// purge cache after a certain interval
func (c *Cache) reap(interval time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	timeAgo := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timeAgo) {
			delete(c.cache, k)
		}
	}
}
