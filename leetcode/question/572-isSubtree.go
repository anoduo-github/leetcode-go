package question

//572. 另一棵树的子树
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil && subRoot != nil {
		return false
	}
	return isSame(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSame(r, s *TreeNode) bool {
	if s == nil && r == nil {
		return true
	}
	if s != nil && r != nil {
		if s.Val == r.Val {
			return isSame(r.Left, s.Left) && isSame(r.Right, s.Right)
		}
	}
	return false
}
