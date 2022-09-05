package question

//687. 最长同值路径
func longestUnivaluePath(root *TreeNode) int {
	res := 0
	var recv func(node *TreeNode) int
	recv = func(node *TreeNode) int {
		//返回条件
		if node == nil {
			return 0
		}
		l := recv(node.Left)
		r := recv(node.Right)
		left := 0  //定义父节点到左子节点的路径长度
		right := 0 //定义父节点到右子节点的路径长度
		//判断左子节点
		if node.Left != nil && node.Left.Val == node.Val {
			left = l + 1
		}
		//判断右子节点
		if node.Right != nil && node.Right.Val == node.Val {
			right = r + 1
		}
		if res < right+left {
			res = right + left
		}

		if right > left {
			return right
		}
		return left
	}
	recv(root)
	return res
}
