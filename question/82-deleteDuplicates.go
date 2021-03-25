package question

//82. 删除排序链表中的重复元素 II
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	//新建一个首节点指向头节点
	first := &ListNode{-101, head}
	//创建一个个临时节点
	cur := first
	//遍历
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			temp := cur.Next.Next
			for temp.Next != nil {
				if temp.Val != temp.Next.Val {
					break
				}
				temp = temp.Next
			}
			cur.Next = temp.Next
		} else {
			cur = cur.Next
		}
	}
	return first.Next
}
