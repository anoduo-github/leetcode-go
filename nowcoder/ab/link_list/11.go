package linklist

//AB12 删除链表的节点
func deleteNode(head *ListNode, val int) *ListNode {
	cur := &ListNode{}
	cur.next = head
	temp := cur
	for cur.next != nil {
		if cur.next.val == val {
			cur.next = cur.next.next
		}
		cur = cur.next
	}
	return temp.next
}
