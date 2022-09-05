package question

type ListNode struct {
	Val  int
	Next *ListNode
}

//92. 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	tempNode := &ListNode{Val: -1}
	tempNode.Next = head
	pre := tempNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return tempNode.Next
}
