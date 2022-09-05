package tree

//AB17 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	// write code here
	if len(inorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{
			Val: inorder[0],
		}
	}
	//递归
	//根据后序遍历特性，最后一个就是根节点
	root := &TreeNode{
		Val: postorder[len(postorder)-1],
	}
	//遍历中序结果
	for i := 0; i < len(inorder); i++ {
		if root.Val == inorder[i] {
			//表示i左边的都是左子节点，右边的都是右子节点
			//递归构建
			root.Left = buildTree(inorder[0:i], postorder[0:i])
			root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
		}
	}
	return root
}
