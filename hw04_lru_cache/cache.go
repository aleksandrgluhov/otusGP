package hw04lrucache

import "sync"

type Key string

// Linked list can receive any value
// So I decided to wrap actual cache value to custom type
// which will store items map key, along with the actual value
// to provide O(1) complexity of Set operation with item eviction

type valueWrapper struct {
	itemsKey   Key
	queueValue interface{}
}

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mu       sync.Mutex
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheItem, ok := c.items[key]
	if ok {
		cacheItem.Value = valueWrapper{itemsKey: key, queueValue: value}
		c.queue.MoveToFront(cacheItem)
		return true
	}
	if c.queue.Len() >= c.capacity {
		delete(c.items, c.queue.Back().Value.(valueWrapper).itemsKey)
		c.queue.Remove(c.queue.Back())
	}
	newCacheItem := c.queue.PushFront(valueWrapper{itemsKey: key, queueValue: value})
	c.items[key] = newCacheItem
	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheItem, ok := c.items[key]
	if ok {
		c.queue.MoveToFront(cacheItem)
		return cacheItem.Value.(valueWrapper).queueValue, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	clear(c.items)
	c.queue = NewList()
}
