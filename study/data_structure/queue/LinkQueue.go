package queue

import (
	"errors"
	"sync"
)

//LinkNode 链表节点
type LinkNode struct {
	value string    //节点值
	next  *LinkNode //下一个节点指针
}

//LinkQueue 链表队列
type LinkQueue struct {
	root *LinkNode  //根节点
	size int        //大小
	lock sync.Mutex //锁
}

//NewLinkQueue 新建一个链表队列
func NewLinkQueue() *LinkQueue {
	return &LinkQueue{
		root: nil,
		size: 0,
		lock: sync.Mutex{},
	}
}

//Add 入队
func (l *LinkQueue) Add(v string) {
	l.lock.Lock()
	defer l.lock.Unlock()

	newNode := &LinkNode{
		value: v,
		next:  nil,
	}

	if l.root == nil {
		l.root = newNode
	} else {
		//遍历节点到末尾
		temp := l.root
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newNode
	}

	l.size++
}

//Remove 出队
func (l *LinkQueue) Remove() (string, error) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if l.size == 0 {
		return "", errors.New("队列中元素已空")
	}

	temp := l.root

	l.root = l.root.next
	l.size--

	return temp.value, nil
}

//Size 队列大小
func (l *LinkQueue) Size() int {
	return l.size
}

//IsEmpty 是否为空
func (l *LinkQueue) IsEmpty() bool {
	return l.size == 0
}
