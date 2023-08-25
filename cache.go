package cacheku

import (
	"sync"
	"time"
)

type Item struct {
	Value  any   `json:"value"`
	Expire int64 `json:"expire"`
}

type Cache struct {
	object     map[string]Item
	defaultExp time.Duration
	mu         sync.RWMutex
	toDelete   chan string
}

func NewCache(defaultExp time.Duration) *Cache {
	c := &Cache{
		object:     make(map[string]Item),
		mu:         sync.RWMutex{},
		defaultExp: defaultExp,
		toDelete:   make(chan string, 10),
	}

	go c.delete()
	return c
}

func (c *Cache) Set(key string, value any, expired time.Duration) error {
	var exp int64

	if expired <= 0 {
		exp = time.Now().Add(c.defaultExp).Unix()
	} else {
		exp = time.Now().Add(expired).Unix()
	}
	c.mu.Lock()
	c.object[key] = Item{
		Value:  value,
		Expire: exp,
	}
	c.mu.Unlock()
	return nil
}

func (c *Cache) Get(key string) (bool, any) {
	c.mu.RLock()
	value, found := c.object[key]
	if !found {
		c.mu.RUnlock()
		return false, nil
	}

	if time.Now().Unix() > value.Expire {
		c.mu.RUnlock()
		return false, nil
	}
	c.mu.RUnlock()
	return true, value.Value
}

func (c *Cache) Delete(key string) {
	c.mu.RLock()
	delete(c.object, key)
	c.mu.RUnlock()
}

func (c *Cache) FetchAll() map[string]Item {
	return c.object
}

func (c *Cache) d() {
	defer close(c.toDelete)
	for key := range c.toDelete {
		c.mu.Lock()
		delete(c.object, key)
		c.mu.Unlock()
	}
}

func (c *Cache) delete() {
	go c.d()

	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		for key, value := range c.object {
			if time.Now().Unix() > value.Expire {
				c.toDelete <- key
			}
		}
	}
}
