package question

//160. 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a, b := headA, headB
	for a != b {
		//循环指向下一个，当下一个为空，则重新指向另一个链表的头，直到两个指针指向同一个，即为第一个交点
		if a != nil {
			a = a.Next
		} else {
			a = headB
		}
		if b != nil {
			b = b.Next
		} else {
			b = headA
		}
	}
	return a
}

//5 6 1 8 4 5 4 1 8 4 5

//4 1 8 4 5 5 6 1 8 4 5
