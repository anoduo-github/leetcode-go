package tree

import "testing"

func Test(t *testing.T) {
	one := &TreeNode{
		Val: 1,
	}
	two := &TreeNode{
		Val: 2,
	}
	three := &TreeNode{
		Val: 3,
	}
	one.Left = two
	one.Right = three
	res := threeOrders(one)
	t.Error(res)
}
