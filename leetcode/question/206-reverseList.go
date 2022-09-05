package question

//206. 反转链表
func reverseList(head *ListNode) *ListNode {
	cur := head        //指针，遍历链表
	res := &ListNode{} //结果，先建一个空的头节点
	for cur != nil {
		temp := cur          //获取当前节点
		cur = cur.Next       //cur指向下一个节点
		temp.Next = res.Next //当前节点的下一个节点指向res的下一个节点
		res.Next = temp      //res的下一个节点就是当前节点
	}
	return res.Next
}
