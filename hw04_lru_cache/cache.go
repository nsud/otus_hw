package hw04_lru_cache //nolint:golint,stylecheck
import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key Key) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                              // Очистить кэш
}
type lruCache struct {
	mx       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}
type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mx.Lock()
	defer c.mx.Unlock()
	val, ok := c.items[key]
	item := cacheItem{key, value}
	//fmt.Printf("Set: %v \t %v \n",val, ok)
	if ok {
		val.Value = item.value
		c.queue.MoveToFront(val)
		return true
	}
	c.queue.PushFront(value)
	c.items[key] = c.queue.Front()
	if c.queue.Len() > c.capacity {
		lastEl := c.queue.Back()
		c.queue.Remove(lastEl)
		//fmt.Printf("Set remove: %v \n", lastEl)
	}
	return false
}
func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mx.Lock()
	defer c.mx.Unlock()
	cItem, ok := c.items[key]
	if ok {
		//fmt.Println(cItem)
		c.queue.MoveToFront(cItem)
		return cItem.Value, ok
	}
	return nil, ok
}
func (c *lruCache) Clear() {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.queue = NewList()
	for key := range c.items {
		delete(c.items, key)
	}
}
func NewCache(capacity int) Cache {
	cache := &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem),
	}
	//fmt.Println(Cache.items)
	return cache
}
