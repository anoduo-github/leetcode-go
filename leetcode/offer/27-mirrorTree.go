package offer

//剑指 Offer 27. 二叉树的镜像
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	if root.Left != nil {
		root.Left = mirrorTree(root.Left)
	}
	if root.Right != nil {
		root.Right = mirrorTree(root.Right)
	}
	return root
}
