package question

import "container/list"

type MyHashSet struct {
	//list 本质是双向链表
	data []list.List
}

/** Initialize your data structure here. */
func Constructor705() MyHashSet {
	return MyHashSet{
		data: make([]list.List, 997),
	}
}

func (this *MyHashSet) hash(key int) int {
	return key % 997
}

func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		k := this.hash(key)
		//PushBack 在末尾插入数据
		this.data[k].PushBack(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	k := this.hash(key)
	//Front 返回链表首元素
	for e := this.data[k].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			//Remove 删除一个链表元素
			this.data[k].Remove(e)
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	k := this.hash(key)
	for e := this.data[k].Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return true
		}
	}
	return false
}
