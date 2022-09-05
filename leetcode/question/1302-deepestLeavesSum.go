package question

var level int = -1
var max int = 0

//1302. 层数最深叶子节点的和
func deepestLeavesSum(root *TreeNode) int {
	//深度遍历
	dfs(root, -1)
	//这么写是为了测试用例重复使用
	res := max
	level = -1
	max = 0
	return res
}

func dfs(node *TreeNode, n int) {
	//为空直接返回
	if node == nil {
		return
	}
	//n表示上一层，n+1表示当前层，相等就相加
	if n+1 == level {
		max += node.Val
	} else if n+1 > level {
		//不等，只可能是当前层大于level,重置
		level = n + 1
		max = node.Val
	}
	//递归
	dfs(node.Left, n+1)
	dfs(node.Right, n+1)
}
