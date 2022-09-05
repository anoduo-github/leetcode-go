package question

//437. 路径总和 III
func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum = sum - node.Val
		if sum == 0 {
			count++
		}
		dfs(node.Left, sum)
		dfs(node.Right, sum)
	}

	if root == nil {
		return 0
	}
	dfs(root, targetSum)
	count += pathSum(root.Left, targetSum)
	count += pathSum(root.Right, targetSum)
	return count
}
