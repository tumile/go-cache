package cache

type lruCache struct {
	list     *doublyLinkedList
	dict     map[interface{}]*node
	capacity int
}

func NewLRUCache(capacity int) *lruCache {
	cache := lruCache{
		list:     newDoublyLinkedList(),
		dict:     map[interface{}]*node{},
		capacity: capacity,
	}
	return &cache
}

func (cache *lruCache) Put(key, val interface{}) {
	if cache.dict[key] != nil {
		cache.list.remove(cache.dict[key])
	} else if cache.list.size == cache.capacity {
		delete(cache.dict, cache.list.removeTail().key)
	}
	n := &node{key, val, nil, nil}
	cache.dict[key] = n
	cache.list.add(n)
}

func (cache *lruCache) Get(key interface{}) interface{} {
	if cache.dict[key] != nil {
		n := cache.dict[key]
		cache.list.remove(n)
		cache.list.add(n)
		return n.val
	}
	return nil
}
