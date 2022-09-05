package linklist

//AB11 合并两个排序的链表
func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	res := &ListNode{}
	cur := res
	for pHead1 != nil && pHead2 != nil {
		if pHead1.val < pHead2.val {
			cur.next = pHead1
			pHead1 = pHead1.next
		} else {
			cur.next = pHead2
			pHead2 = pHead2.next
		}
		cur = cur.next
	}
	if pHead1 != nil {
		cur.next = pHead1
	} else {
		cur.next = pHead2
	}
	return res.next
}
