package linklist

import (
	"fmt"
	"strconv"
)

//AB9 【模板】链表

//ListNode 链表，带头结点
type ListNode struct {
	val  int
	next *ListNode
}

func NewListNode() *ListNode {
	return &ListNode{}
}

func (l *ListNode) Insert(x, y int) {
	cur := l
	node := &ListNode{
		val: y,
	}
	for {
		//未找到
		if cur.next == nil {
			cur.next = node
			break
		}
		//找到匹配的值
		if cur.next.val == x {
			temp := cur.next
			cur.next = node
			node.next = temp
			break
		}
		cur = cur.next
	}
}

func (l *ListNode) Delete(x int) {
	cur := l
	for {
		//未找到
		if cur.next == nil {
			break
		}
		//找到匹配的值
		if cur.next.val == x {
			cur.next = cur.next.next
			break
		}
		cur = cur.next
	}
}

func (l *ListNode) Print() {
	cur := l
	str := ""
	for {
		//未找到
		if cur.next == nil {
			break
		}
		str += strconv.Itoa(cur.next.val) + " "
		cur = cur.next
	}
	if len(str) == 0 {
		fmt.Println("NULL")
	} else {
		fmt.Println(str)
	}
}
