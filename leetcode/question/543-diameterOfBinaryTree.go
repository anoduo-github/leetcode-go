package question

//543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	max := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		//递归获取左右子节点的值
		left := dfs(node.Left)
		right := dfs(node.Right)
		//判断当前节点与子节点的最大长度
		if node.Left != nil {
			left++
		}
		if node.Right != nil {
			right++
		}
		max = maxxx(max, left+right)
		return maxxx(left, right)
	}
	dfs(root)
	return max
}

func maxxx(a, b int) int {
	if a > b {
		return a
	}
	return b
}
