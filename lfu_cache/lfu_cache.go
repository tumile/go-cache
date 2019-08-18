package lfu_cache

import "log"

type LFUCache struct {
	freqList []*doublyLinkedList
	dict     map[interface{}]*node
	size     int
	capacity int
	minFreq  int
}

func NewLFUCache(capacity int) *LFUCache {
	if capacity == 0 {
		log.Fatal("Capacity should not be 0")
	}
	freqList := make([]*doublyLinkedList, capacity)
	for i := 0; i < capacity; i++ {
		freqList[i] = newDoublyLinkedList()
	}
	cache := LFUCache{
		freqList: freqList,
		dict:     map[interface{}]*node{},
		size:     0,
		capacity: capacity,
		minFreq:  0,
	}
	return &cache
}

func (cache *LFUCache) Get(key interface{}) interface{} {
	if _, ok := cache.dict[key]; ok {
		node := cache.dict[key]
		cache.moveToNextBucket(node)
		return node.val
	}
	return -1
}

func (cache *LFUCache) Put(key, val interface{}) {
	if _, ok := cache.dict[key]; ok {
		node := cache.dict[key]
		node.val = val
		cache.moveToNextBucket(node)
		return
	}
	if cache.size == cache.capacity {
		minFreqNode := cache.freqList[cache.minFreq].removeTail()
		delete(cache.dict, minFreqNode.key)
		cache.size--
	}
	node := &node{key, val, 0, nil, nil}
	cache.freqList[0].add(node)
	cache.dict[key] = node
	cache.minFreq = 0
	cache.size++
}

func (cache *LFUCache) moveToNextBucket(node *node) {
	cache.freqList[node.bucket].remove(node)
	if node.bucket < cache.capacity-1 {
		if node.bucket == cache.minFreq && cache.freqList[node.bucket].size == 0 {
			cache.minFreq++
		}
		node.bucket++
	}
	cache.freqList[node.bucket].add(node)
}
