package offer

//剑指 Offer 18. 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	temp := &ListNode{
		Val:  0,
		Next: head,
	}
	cur := temp
	for cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			break
		}
		cur = cur.Next
	}
	return temp.Next
}
