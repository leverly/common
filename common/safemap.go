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

// all kv pair count
func (this *SafeMap) Len() int {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return len(this.container)
}

// if key exist update, else insert
func (this *SafeMap) Replace(key interface{}, value interface{}) (interface{}, bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	temp, find := this.container[key]
	this.container[key] = value
	return temp, find
}

// if key exist return err, else return nil
func (this *SafeMap) Insert(key interface{}, value interface{}) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, find := this.container[key]
	if find {
		return ErrEntryExist
	}
	this.container[key] = value
	return nil
}

// if key exist return nil, else return err
func (this *SafeMap) Update(key interface{}, value interface{}) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	_, find := this.container[key]
	if !find {
		return ErrEntryNotExist
	}
	this.container[key] = value
	return nil
}

// if key exist return value + true, else return nil + false
func (this *SafeMap) Delete(key interface{}) (interface{}, bool) {
	this.lock.Lock()
	defer this.lock.Unlock()
	value, find := this.container[key]
	delete(this.container, key)
	return value, find
}

// if key exist return value + true, else return nil + false
func (this *SafeMap) Find(key interface{}) (interface{}, bool) {
	this.lock.RLock()
	defer this.lock.RUnlock()
	value, find := this.container[key]
	return value, find
}
