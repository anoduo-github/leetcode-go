package offer

//剑指 Offer 26. 树的子结构
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	//这里只要满足其中一个就行
	//即要么a,b都是从根节点开始，isSub(A,B)
	//要么就是b是a的左右节点子树,isSubStructure() 递归处理左右子节点
	return isSub(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSub(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		//表示b已经遍历完了
		return true
	}
	if a == nil || a.Val != b.Val {
		//表示a遍历完了，b没完；或者a的节点值与b不等
		return false
	}
	//继续比较左右子节点
	return isSub(a.Left, b.Left) && isSub(a.Right, b.Right)
}
