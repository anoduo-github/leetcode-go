package question

//142. 环形链表 II
func detectCycle(head *ListNode) *ListNode {
	m := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := m[head]; ok {
			return head
		}
		m[head] = struct{}{}
		head = head.Next
	}
	return nil
}
