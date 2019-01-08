package lru

import (
	"container/list"
	"errors"
)

// LRU implements a fixed size LRU cache
type LRU struct {
	size      int
	cacheList *list.List
	items     map[interface{}]*list.Element
}

// item defines the item in the LRU list
type cacheItem struct {
	key   interface{}
	value interface{}
}

// NewLRU constructs an LRU cache
func NewLRU(size int) (*LRU, error) {
	if size <= 0 {
		return nil, errors.New("invalid size")
	}
	c := &LRU{
		size:      size,
		cacheList: list.New(),
		items:     make(map[interface{}]*list.Element),
	}
	return c, nil
}

// Clear clear all the items in LRU cache
func (c *LRU) Clear() {
	// delete map item
	for k := range c.items {
		delete(c.items, k)
	}
	// clear list
	c.cacheList.Init()
}

// Put add item to LRU cache, if the key already exists, update the value
func (c *LRU) Put(key, value interface{}) {
	// check if the key exists
	if item, ok := c.items[key]; ok {
		// move the item to the front
		c.cacheList.MoveToFront(item)
		// update the item value
		item.Value.(*cacheItem).value = value
		return
	}

	// add new item
	item := &cacheItem{key, value}
	item1 := c.cacheList.PushFront(item)
	c.items[key] = item1

	//check if the cache is full
	if c.cacheList.Len() > c.size {
		c.removeOldest()
	}
}

// Get return the value of the key, nil if not exists
func (c *LRU) Get(key interface{}) (value interface{}, err error) {
	if item, ok := c.items[key]; ok {
		c.cacheList.MoveToFront(item)
		return item.Value.(*cacheItem).value, nil
	}
	return nil, errors.New("empty")
}

func (c *LRU) removeOldest() {
	item := c.cacheList.Back()
	if item != nil {
		c.removeItem(item)
	}
}

// removeItem . remove item in the LRU cache
func (c *LRU) removeItem(item *list.Element) {
	// remove item from list
	c.cacheList.Remove(item)
	kv := item.Value.(*cacheItem)
	// delete key in the map
	delete(c.items, kv.key)
}

// Del . delete the key in LRU cache, return true if exists, else return false
func (c *LRU) Del(key interface{}) bool {
	if item, ok := c.items[key]; ok {
		c.removeItem(item)
		return true
	}
	return false
}

// Len . return the number of items in the cache
func (c *LRU) Len() int {
	return c.cacheList.Len()
}
