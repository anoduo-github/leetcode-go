package question

// 965. 单值二叉树
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	temp := root.Val
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != temp {
			return false
		}
		return dfs(node.Left) && dfs(node.Right)
	}
	return dfs(root.Left) && dfs(root.Right)
}
