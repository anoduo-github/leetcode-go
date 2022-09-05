package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) indexOrder(res *[]int) {
	*res = append(*res, t.Val)
	if t.Left != nil {
		t.Left.indexOrder(res)
	}
	if t.Right != nil {
		t.Right.indexOrder(res)
	}
}

func (t *TreeNode) IndexOrder() []int {
	res := make([]int, 0)
	t.indexOrder(&res)
	return res
}

func (t *TreeNode) midOrder(res *[]int) {
	if t.Left != nil {
		t.Left.midOrder(res)
	}
	*res = append(*res, t.Val)
	if t.Right != nil {
		t.Right.midOrder(res)
	}
}

func (t *TreeNode) MidOrder() []int {
	res := make([]int, 0)
	t.midOrder(&res)
	return res
}

func (t *TreeNode) lastOrder(res *[]int) {
	if t.Left != nil {
		t.Left.lastOrder(res)
	}
	if t.Right != nil {
		t.Right.lastOrder(res)
	}
	*res = append(*res, t.Val)
}

func (t *TreeNode) LastOrder() []int {
	res := make([]int, 0)
	t.lastOrder(&res)
	return res
}

/**
 * AB16 实现二叉树先序，中序和后序遍历
 * @param root TreeNode类 the root of binary tree
 * @return int整型二维数组
 */
func threeOrders(root *TreeNode) [][]int {
	// write code here
	res := make([][]int, 3, 3)
	if root == nil {
		return res
	}
	index := root.IndexOrder()
	res[0] = index
	mid := root.MidOrder()
	res[1] = mid
	last := root.LastOrder()
	res[2] = last
	return res
}
