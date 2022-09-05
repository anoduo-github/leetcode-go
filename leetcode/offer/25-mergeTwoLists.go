package offer

//剑指 Offer 25. 合并两个排序的链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{
		Val: 0,
	}
	cur := res
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
		} else {
			cur.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return res.Next
}
