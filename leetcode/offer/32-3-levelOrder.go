package offer

//剑指 Offer 32 - III. 从上到下打印二叉树 III
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)       //结果集
	level := 1                    //第一层
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
		if level%2 == 0 {
			revs(vals)
		}
		res = append(res, vals)
		queue = temp
		level++
	}
	return res
}

func revs(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
