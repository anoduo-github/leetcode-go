package question

//814. 二叉树剪枝
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		root.Left = pruneTree(root.Left)
	}
	if root.Right != nil {
		root.Right = pruneTree(root.Right)
	}
	if root.Val == 0 && root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}
