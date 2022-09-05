package question

type node struct {
	Val      int
	Children []*node
}

//429. N 叉树的层序遍历
func levelOrder(root *node) [][]int {
	//初始化二维数组
	res := make([][]int, 0)
	for i := range res {
		res[i] = make([]int, 0)
	}
	//为空直接返回
	if root == nil {
		return res
	}
	//用队列存放存在子树的节点
	queue := make([]*node, 0)
	queue = append(queue, root)
	//遍历queue
	for len(queue) > 0 {
		temp := make([]*node, 0) //下一层节点
		rows := make([]int, 0)   //结果
		for i := 0; i < len(queue); i++ {
			//存值
			rows = append(rows, queue[i].Val)
			//表示有子节点
			if len(queue[i].Children) > 0 {
				//扩展队列
				temp = append(temp, queue[i].Children...)
			}
		}
		//替换
		queue = temp
		res = append(res, rows)
	}
	return res
}
