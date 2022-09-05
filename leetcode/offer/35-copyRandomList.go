package offer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

//剑指 Offer 35. 复杂链表的复制
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	//先复制一份
	for cur := head; cur != nil; cur = cur.Next.Next {
		temp := cur.Next
		cur.Next = &Node{
			Val:  cur.Val,
			Next: temp,
		}
	}
	//在复制随机指针
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	//拆分成两个
	res := head.Next
	for cur := head; cur != nil; cur = cur.Next {
		temp := cur.Next
		cur.Next = cur.Next.Next
		if temp.Next != nil {
			temp.Next = temp.Next.Next
		}
	}
	return res
}
