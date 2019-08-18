package lru_cache

import "log"

type LRUCache struct {
	list     *doublyLinkedList
	dict     map[interface{}]*node
	capacity int
}

func NewLRUCache(capacity int) *LRUCache {
	if capacity == 0 {
		log.Fatal("Capacity should not be 0")
	}
	cache := LRUCache{
		list:     newDoublyLinkedList(),
		dict:     map[interface{}]*node{},
		capacity: capacity,
	}
	return &cache
}

func (cache *LRUCache) Get(key interface{}) interface{} {
	if _, ok := cache.dict[key]; ok {
		node := cache.dict[key]
		cache.list.remove(node)
		cache.list.add(node)
		return node.val
	}
	return nil
}

func (cache *LRUCache) Put(key, val interface{}) {
	node := &node{key, val, nil, nil}
	if _, ok := cache.dict[key]; ok {
		cache.list.remove(cache.dict[key])
	} else if cache.list.size == cache.capacity {
		k := cache.list.removeTail().key
		delete(cache.dict, k)
	}
	cache.list.add(node)
	cache.dict[key] = node
}
