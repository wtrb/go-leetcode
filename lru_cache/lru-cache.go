package cache

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	lru      *list.List
}

type keyVal struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		lru:      list.New(),
	}
}

// Time complexity: O(1)
func (this *LRUCache) Get(key int) int {
	if elem, ok := this.cache[key]; ok {
		this.lru.MoveToFront(elem)
		return elem.Value.(*keyVal).val

	} else {
		return -1
	}
}

// Time complexity: O(1)
func (this *LRUCache) Put(key int, value int) {
	if elem, ok := this.cache[key]; ok {
		this.lru.MoveToFront(elem)
		elem.Value = &keyVal{key, value}

	} else {
		if len(this.cache) == this.capacity {
			back := this.lru.Back()
			delete(this.cache, back.Value.(*keyVal).key)
			this.lru.Remove(back)
		}

		this.lru.PushFront(&keyVal{key, value})
		this.cache[key] = this.lru.Front()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// LRU Cache

// https://leetcode.com/problems/lru-cache/
// https://leetcode.com/submissions/detail/329410774/?from=/explore/featured/card/30-day-leetcoding-challenge/531/week-4/3309/

// tags: lru, cache, least recently used
