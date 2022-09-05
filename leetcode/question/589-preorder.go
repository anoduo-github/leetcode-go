package question

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

//589. N 叉树的前序遍历
func preorder(root *Node) []int {
	return pre(root)
}

//pre 前序遍历
func pre(root *Node) []int {
	res := make([]int, 0)
	//为空直接返回
	if root == nil {
		return res
	}
	//存值
	res = append(res, root.Val)
	//遍历子节点
	for _, v := range root.Children {
		//递归
		temp := pre(v)
		res = append(res, temp...)
	}
	return res
}

//使用栈实现
func preorder2(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	//使用切片代替栈
	stack := make([]*Node, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		temp := stack[len(stack)-1]
		res = append(res, temp.Val)
		//让temp出栈
		stack = stack[:len(stack)-1]
		//从右至左遍历temp子节点，入栈
		for i := len(temp.Children) - 1; i >= 0; i-- {
			stack = append(stack, temp.Children[i])
		}
	}

	return res
}
