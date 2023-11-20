package main

import (
	"container/list"
	"fmt"
)

type LRUCache[T any] interface {
	Get(key int) (T, bool)
	Put(key int, value T)
}

type LRUCacheImpl[T any] struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Node[T any] struct {
	key   int
	value T
}

func NewLRUCache[T any](capacity int) *LRUCacheImpl[T] {
	return &LRUCacheImpl[T]{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (lru *LRUCacheImpl[T]) Get(key int) (T, bool) {
	var zero T
	if elem, ok := lru.cache[key]; ok {
		lru.list.MoveToFront(elem)
		return elem.Value.(*Node[T]).value, true
	}
	return zero, false
}

func (lru *LRUCacheImpl[T]) Put(key int, value T) {
	if elem, ok := lru.cache[key]; ok {
		elem.Value.(*Node[T]).value = value
		lru.list.MoveToFront(elem)
	} else {
		if len(lru.cache) >= lru.capacity {
			back := lru.list.Back()
			delete(lru.cache, back.Value.(*Node[T]).key)
			lru.list.Remove(back)
		}
		newNode := &Node[T]{key, value}
		elem := lru.list.PushFront(newNode)
		lru.cache[key] = elem
	}
}

func main() {
	cache := NewLRUCache[int](2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(2)) // Output: 1, true

	cache.Put(3, 3)
	fmt.Println(cache.Get(1)) // Output: false
	cache.Put(4, 4)
	fmt.Println(cache.Get(1)) // Output: false
	fmt.Println(cache.Get(3)) // Output: 3
	fmt.Println(cache.Get(4)) // Output: 4
}
