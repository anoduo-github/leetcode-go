package question

//61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	//特殊情况下
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	//将链表改为环形
	cur := head
	length := 1
	for cur.Next != nil {
		cur = cur.Next
		length++
	}
	//表示旋转一周回到原点
	if k%length == 0 {
		return head
	}
	//成环
	cur.Next = head
	//移动cur,找到新头节点
	count := length - k%length
	for count > 0 {
		cur = cur.Next
		count--
	}
	res := cur.Next
	cur.Next = nil
	return res
}
