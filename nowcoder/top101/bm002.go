package top101

// BM2 链表内指定区间反转
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	temp := new(ListNode)
	temp.Next = head
	pre := temp
	for i := 0; i < m-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	var next *ListNode
	for i := 0; i < n-m; i++ {
		next = cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return temp.Next
}
