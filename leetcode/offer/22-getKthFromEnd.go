package offer

//剑指 Offer 22. 链表中倒数第k个节点
func getKthFromEnd(head *ListNode, k int) *ListNode {
	//快慢指针
	l := head
	f := head
	for k > 1 {
		if l != nil {
			l = l.Next
			k--
		} else {
			return f
		}
	}
	for f != nil && l != nil {
		if l.Next == nil {
			return f
		}
		f = f.Next
		l = l.Next
	}
	return f
}
