package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//94. 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return infix(root)
}

//infix 中序遍历
func infix(root *TreeNode) []int {
	res := make([]int, 0)
	//先遍历左节点
	if root.Left != nil {
		left := infix(root.Left)
		res = append(res, left...)
	}
	//再输出父节点
	res = append(res, root.Val)
	//再遍历右节点
	if root.Right != nil {
		right := infix(root.Right)
		res = append(res, right...)
	}
	return res
}
