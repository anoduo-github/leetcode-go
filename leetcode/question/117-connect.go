package question

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

//117. 填充每个节点的下一个右侧节点指针 II
func connect(root *Node117) *Node117 {
	if root == nil {
		return nil
	}
	//使用队列，逐层构建
	queue := make([]*Node117, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		//创建一个临时队列，存遍历的节点
		temp := make([]*Node117, 0)
		//遍历当前队列里的节点
		for i := 0; i < len(queue); i++ {
			if i+1 < len(queue) {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				temp = append(temp, queue[i].Left)
			}
			if queue[i].Right != nil {
				temp = append(temp, queue[i].Right)
			}
		}
		//替换队列
		queue = temp
	}
	return root
}
