package question

//226. 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		//左右子节点交换
		node.Left, node.Right = node.Right, node.Left
		//递归左右子节点
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)

	return root
}
