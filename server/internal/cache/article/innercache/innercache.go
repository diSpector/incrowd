package innercache

import (
	"context"
	"log"
	"sync"
	"time"
)

type Cache struct {
	mx    sync.RWMutex
	cache map[string]value
	ttl   time.Duration
}

type value struct {
	val    string
	expire time.Time
}

func New(ctx context.Context, ttl time.Duration) *Cache {
	cache := &Cache{
		cache: make(map[string]value),
		ttl:   ttl,
	}

	go cache.invalidateLoop(ctx)

	return cache
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	if v, ok := c.cache[key]; ok {
		v.expire = time.Now().Add(c.ttl)
		c.cache[key] = v
		return v.val, nil
	}

	return ``, ErrNotFoundInCache
}

func (c *Cache) Set(ctx context.Context, key, val string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.cache[key] = value{
		val:    val,
		expire: time.Now().Add(c.ttl),
	}

	return nil
}

func (c *Cache) delete(key string) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.cache, key)

	return nil
}

func (c *Cache) invalidateLoop(ctx context.Context) {
	for {
		for k, v := range c.cache {
			if time.Now().After(v.expire) {
				c.delete(k)
			}
		}
		select {
		case <-ctx.Done():
			log.Println(`cache invalidate loop stopped`)
			return
		case <-time.After(10 * time.Minute):
		}
	}
}
