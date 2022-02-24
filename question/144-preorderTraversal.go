package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//144.二叉树的前序遍历
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	return preOrder(root)
}

//preOrder 前序遍历
func preOrder(root *TreeNode) []int {
	res := make([]int, 0)
	//首先收集父节点
	res = append(res, root.Val)
	//再遍历左节点
	if root.Left != nil {
		left := preOrder(root.Left)
		res = append(res, left...)
	}
	//最后遍历右节点
	if root.Right != nil {
		right := preOrder(root.Right)
		res = append(res, right...)
	}
	return res
}
