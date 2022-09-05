package question

import "testing"

func TestPruneTree(t *testing.T) {
	root := &TreeNode{
		Val: 1,
	}
	right := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	root.Right = right
	res := pruneTree(root)
	t.Log(res)
}
