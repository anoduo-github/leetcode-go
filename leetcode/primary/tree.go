package primary

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//LC 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1)
}

//全局前一个点
var prev *TreeNode

//LC 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	prev = nil
	return isBst(root)
}

func isBst(root *TreeNode) bool {
	//中序遍历
	if root == nil {
		return true
	}
	//左节点
	if !isBst(root.Left) {
		return false
	}
	//中间节点
	if prev != nil && prev.Val >= root.Val {
		return false
	}
	//否则重新指向
	prev = root
	//右节点
	return isBst(root.Right)
}

//LC 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricNode(root.Left, root.Right)
}

func isSymmetricNode(left, right *TreeNode) bool {
	//当左右子节点都为空
	if left == nil && right == nil {
		return true
	}
	//当左右子节点不都为空且值不同
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	//递归判断左子节点的左子节点和右子节点的右子节点，左子节点的右子节点和右子节点的左子节点
	return isSymmetricNode(left.Left, right.Right) && isSymmetricNode(left.Right, right.Left)
}

//LC 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	//bfs
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	//建立一个切片存每层节点
	queue := make([]*TreeNode, 0)
	//首先是第一层根节点root
	queue = append(queue, root)
	for len(queue) > 0 {
		order := make([]int, 0)      //存每层结果
		temp := make([]*TreeNode, 0) //存下一层节点
		//遍历queue
		for _, v := range queue {
			order = append(order, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		res = append(res, order)
		//重新赋值queue
		queue = temp
	}
	return res
}

//LC 将有序数组转换为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return sortedArrayToNode(nums, 0, len(nums)-1)
}

func sortedArrayToNode(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	//取中间值
	mid := (start + end) / 2
	//创建节点
	node := &TreeNode{
		Val: nums[mid],
	}
	node.Left = sortedArrayToNode(nums, start, mid-1)
	node.Right = sortedArrayToNode(nums, mid+1, end)
	return node
}
