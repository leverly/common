package common

import (
	"sync"
)

type SafeMap struct {
	lock      sync.RWMutex
	container map[interface{}](interface{})
}

func NewSafeMap() *SafeMap {
	conn := make(map[interface{}](interface{}))
	return &SafeMap{container: conn}
}

func (this *SafeMap) Len() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.container)
}

// key must not exist
func (this *SafeMap) Insert(key interface{}, value interface{}) (err error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, find := this.container[key]
	if find {
		return ErrEntryExist
	}
	this.container[key] = value
	return err
}

func (this *SafeMap) Replace(key interface{}, value interface{}) (interface{}, bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	temp, find := this.container[key]
	this.container[key] = value
	return temp, find
}

func (this *SafeMap) Update(key interface{}, value interface{}) (err error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, find := this.container[key]
	if !find {
		return ErrEntryNotExist
	}
	this.container[key] = value
	return err
}

func (this *SafeMap) Delete(key interface{}) (interface{}, bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	value, find := this.container[key]
	delete(this.container, key)
	return value, find
}

func (this *SafeMap) Find(key interface{}) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	value, find := this.container[key]
	return value, find
}
