package offer

//剑指 Offer 32 - II. 从上到下打印二叉树 II
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)       //结果集
	queue := make([]*TreeNode, 0) //节点队列
	queue = append(queue, root)   //添加第一个
	for len(queue) > 0 {          //不为空时
		//存临时节点
		temp := make([]*TreeNode, 0)
		//存结果
		vals := make([]int, 0)
		//遍历队列节点
		for _, v := range queue {
			vals = append(vals, v.Val)
			if v.Left != nil {
				temp = append(temp, v.Left)
			}
			if v.Right != nil {
				temp = append(temp, v.Right)
			}
		}
		res = append(res, vals)
		queue = temp
	}
	return res
}
