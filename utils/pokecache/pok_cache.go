package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	cache map[string]CacheEntry
	mux *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration)Cache {
	cac:=Cache{
		cache:make(map[string]CacheEntry),
		mux: &sync.Mutex{},
	}
	go cac.reaper(interval)
	return cac
}

func (c *Cache)Add(key string,val []byte){
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key]=CacheEntry{
		createdAt:time.Now(),
		val:val,
	}
}
func(c *Cache)Get(key string)([]byte,bool){
	c.mux.Lock()
	defer c.mux.Unlock()
	entry,ok:=c.cache[key]
	if !ok{
		return nil,false
	}
	return entry.val,true
}
func (c *Cache)reaper(interval time.Duration){
	ticker:=time.NewTicker(interval)
	for range ticker.C{
		c.reap(time.Now().UTC(),interval)
	}
}
func (c *Cache)reap(now time.Time,interval time.Duration){
	c.mux.Lock()
	defer c.mux.Unlock()
	for key,entry:=range c.cache{
		if entry.createdAt.Before(now.Add(-interval)){
			delete(c.cache,key)
		}
	}
}