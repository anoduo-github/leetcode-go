package primary

import "testing"

func TestIsVaild(t *testing.T) {
	node := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
	}
	res := isValidBST(node)
	t.Log(res)
}
