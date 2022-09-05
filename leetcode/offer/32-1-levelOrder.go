package offer

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)         //结果集
	queue := make([]*TreeNode, 0) //节点队列
	queue = append(queue, root)   //添加第一个
	for len(queue) > 0 {          //不为空时
		temp := queue[0]  //取第一个
		queue = queue[1:] //调整队列
		res = append(res, temp.Val)
		if temp.Left != nil {
			queue = append(queue, temp.Left)
		}
		if temp.Right != nil {
			queue = append(queue, temp.Right)
		}
	}
	return res
}
