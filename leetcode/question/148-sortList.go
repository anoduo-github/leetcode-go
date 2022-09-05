package question

import "sort"

//148. 排序链表
func sortList(head *ListNode) *ListNode {
	temp := make([]int, 0)
	cur1, cur2 := head, head
	for cur1 != nil {
		temp = append(temp, cur1.Val)
		cur1 = cur1.Next
	}
	sort.Ints(temp)
	i := 0
	for cur2 != nil && i < len(temp) {
		cur2.Val = temp[i]
		cur2 = cur2.Next
		i++
	}
	return head
}
