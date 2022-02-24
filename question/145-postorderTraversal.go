package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//145. 二叉树的后序遍历
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return postOrder(root)
}

//postOrder 后序遍历
func postOrder(root *TreeNode) []int {
	res := make([]int, 0)
	//先遍历左节点
	if root.Left != nil {
		left := postOrder(root.Left)
		res = append(res, left...)
	}
	//再遍历右节点
	if root.Right != nil {
		right := postOrder(root.Right)
		res = append(res, right...)
	}
	//最后添加本节点
	res = append(res, root.Val)
	return res
}
