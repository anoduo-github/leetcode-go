package question

import "container/list"

type MyHashMap struct {
	data []list.List
}

type entry struct {
	key   int
	value int
}

/** Initialize your data structure here. */
func Constructor2() MyHashMap {
	return MyHashMap{
		data: make([]list.List, 997),
	}
}

func (this *MyHashMap) hash(key int) int {
	return key % 997
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	k := this.hash(key)
	for e := this.data[k].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			e.Value = entry{key, value}
			return
		}
	}
	this.data[k].PushBack(entry{key, value})
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	k := this.hash(key)
	for e := this.data[k].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			return e.Value.(entry).value
		}
	}
	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	k := this.hash(key)
	for e := this.data[k].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			this.data[k].Remove(e)
		}
	}
}
