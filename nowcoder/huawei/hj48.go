package huawei

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

//HJ48 从单向链表中删除指定值的节点
func HJ48(msg string) {
	msg = strings.TrimSpace(msg)
	arr := strings.Split(msg, " ")
	//头节点
	val, _ := strconv.Atoi(arr[1])
	head := &Node{
		Val: val,
	}
	//遍历获取节点
	for i := 2; i < len(arr)-2; i = i + 2 {
		cur := head
		next := cur.Next
		a, _ := strconv.Atoi(arr[i])
		b, _ := strconv.Atoi(arr[i+1])
		for cur != nil {
			if cur.Val == b {
				cur.Next = &Node{
					Val:  a,
					Next: next,
				}
				break
			}
			cur = cur.Next
			next = cur.Next
		}
	}

	target, _ := strconv.Atoi(arr[len(arr)-1])
	cur := &Node{
		Val:  -1,
		Next: head,
	}
	next := cur.Next
	temp := cur
	for next != nil {
		if next.Val == target {
			cur.Next = next.Next
			break
		}
		cur = cur.Next
		next = cur.Next
	}
	temp = temp.Next
	for temp != nil {
		fmt.Print(temp.Val)
		fmt.Print(" ")
		temp = temp.Next
	}
}
