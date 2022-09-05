package stack

import (
	"errors"
	"sync"
)

//LinkNode 链表节点
type LinkNode struct {
	value string    //节点值
	next  *LinkNode //下一个节点指针
}

//LinkStack 链表栈
type LinkStack struct {
	root *LinkNode  //根节点
	size int        //大小
	lock sync.Mutex //锁
}

//NewLinkStack 新建一个链表栈
func NewLinkStack() *LinkStack {
	return &LinkStack{
		root: nil,
		size: 0,
		lock: sync.Mutex{},
	}
}

//Push 入栈
func (ls *LinkStack) Push(v string) {
	//注意:根节点就相当于栈顶，每次push后，根节点里的值永远是最后加入的那个

	ls.lock.Lock()         //加锁
	defer ls.lock.Unlock() //解锁

	//新建节点
	newNode := &LinkNode{
		value: v,
		next:  nil,
	}

	//如果根节点为空,则新建一个节点
	if ls.root == nil {
		ls.root = newNode
	} else {
		temp := ls.root     //备份之前的根节点
		ls.root = newNode   //将根节点指向新建节点
		ls.root.next = temp //根节点下一指针指向备份节点
	}
	ls.size++
}

//Pop 出栈
func (ls *LinkStack) Pop() (string, error) {
	ls.lock.Lock()
	defer ls.lock.Unlock()

	if ls.size == 0 {
		return "", errors.New("栈中元素已空")
	}

	temp := ls.root //备份根节点

	ls.root = temp.next //将根节点指向下一个
	ls.size--

	return temp.value, nil
}

//Peek 获取栈顶元素值
func (ls *LinkStack) Peek() (string, error) {
	if ls.size == 0 {
		return "", errors.New("栈中元素已空")
	}
	return ls.root.value, nil
}

//Size 栈大小
func (ls *LinkStack) Size() int {
	return ls.size
}

//IsEmpty 是否为空
func (ls *LinkStack) IsEmpty() bool {
	return ls.size == 0
}
