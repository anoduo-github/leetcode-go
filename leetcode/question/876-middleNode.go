package question

//876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head
	for q != nil && q.Next != nil {
		q = q.Next.Next
		p = p.Next
	}
	return p
}
