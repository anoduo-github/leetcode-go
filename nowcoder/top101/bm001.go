package top101

type ListNode struct {
	Val  int
	Next *ListNode
}

// BM101 反转链表
func reverseList(pHead *ListNode) *ListNode {
	if pHead == nil {
		return nil
	}
	var pre *ListNode
	cur := pHead
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
