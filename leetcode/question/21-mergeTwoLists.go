package question

//21. 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	//先建一个头节点
	head := &ListNode{
		Val: 0,
	}
	cur := head
	for list1 != nil && list2 != nil {
		var temp *ListNode
		if list1.Val < list2.Val {
			temp = list1
			list1 = list1.Next
		} else {
			temp = list2
			list2 = list2.Next
		}
		cur.Next = temp
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return head.Next
}
