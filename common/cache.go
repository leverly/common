package common

import (
	"container/list"
	"fmt"
)

type Cache interface {
	Get(key interface{}) (interface{}, bool)
	Set(key interface{}, value interface{})
	Delete(key interface{})
	Clear()
}

type LRUCacheItem struct {
	value   interface{}
	element *list.Element
}

func NewLRUCacheItem(value interface{}, element *list.Element) *LRUCacheItem {
	return &LRUCacheItem{value: value, element: element}
}

type LRUCache struct {
	index    map[interface{}](*LRUCacheItem)
	lru      *list.List
	maxCount int64
}

func NewLRUCache(count int64) *LRUCache {
	return &LRUCache{index: make(map[interface{}](*LRUCacheItem)), lru: list.New(), maxCount: count}
}

func (this *LRUCache) Len() int64 {
	return int64(len(this.index))
}

func (this *LRUCache) Get(key interface{}) (interface{}, bool) {
	item, find := this.index[key]
	if find {
		this.lru.MoveToFront(item.element)
		return item.value, true
	}
	return nil, find
}

func (this *LRUCache) Set(key interface{}, value interface{}) {
	item, find := this.index[key]
	if find {
		item.value = value
		this.lru.MoveToFront(item.element)
		return
	} else if this.Len() >= this.maxCount {
		delete(this.index, this.lru.Remove(this.lru.Back()))
	}
	this.index[key] = NewLRUCacheItem(value, this.lru.PushFront(key))
}

func (this *LRUCache) Delete(key interface{}) {
	item, find := this.index[key]
	if find {
		this.lru.Remove(item.element)
		delete(this.index, key)
	}
}

func (this *LRUCache) Clear() {
	this.index = make(map[interface{}](*LRUCacheItem))
	this.lru = list.New()
}

func (this *LRUCache) Debug() {
	fmt.Printf("cache length:<%d %d>, cache list:", this.lru.Len(), len(this.index))
	element := this.lru.Front()
	for element != nil {
		value, find := this.index[element.Value]
		if !find {
			panic("error")
		}
		fmt.Print("<", element.Value, value.value, ">")
		element = element.Next()
	}
	fmt.Println()
}
