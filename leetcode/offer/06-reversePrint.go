package offer

type ListNode struct {
	Val  int
	Next *ListNode
}

//剑指 Offer 06. 从尾到头打印链表
func reversePrint(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	//反转数组
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}
