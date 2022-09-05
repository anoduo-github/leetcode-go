package clue

import "fmt"

//heroNode 节点
type heroNode struct {
	no      int       //编号
	name    string    //名称
	left    *heroNode //左指针
	right   *heroNode //右指针
	noLeft  int       //0表示指向左子树，1表示指向前驱节点
	noRight int       //0表示指向右子树，1表示指向后继节点
	parent  *heroNode //父节点，后序线索化遍历需要用
}

//delNode 删除一个子节点
func (h *heroNode) delNode(no int) {
	//如果当前节点的左节点不为空，则判断结点值是否相等，相等则删除
	if h.left != nil && h.left.no == no {
		h.left = nil
	}

	//如果当前节点的右节点不为空，则判断结点值是否相等，相等则删除
	if h.right != nil && h.right.no == no {
		h.right = nil
	}

	//否则递归
	if h.left != nil {
		h.left.delNode(no)
	}
	if h.right != nil {
		h.right.delNode(no)
	}
}

//preOrder 前序遍历
func (h *heroNode) preOrder() {
	//首先输出根节点
	fmt.Printf("no=%d, name=%s\n", h.no, h.name)
	//递归左
	if h.left != nil {
		h.left.preOrder()
	}
	//递归右
	if h.right != nil {
		h.right.preOrder()
	}
}

//midOrder 中序遍历
func (h *heroNode) midOrder() {
	//递归左
	if h.left != nil {
		h.left.midOrder()
	}

	//输出根节点
	fmt.Printf("no=%d, name=%s\n", h.no, h.name)

	//递归右
	if h.right != nil {
		h.right.midOrder()
	}
}

//postOrder 后序遍历
func (h *heroNode) postOrder() {
	//递归左
	if h.left != nil {
		h.left.postOrder()
	}

	//递归右
	if h.right != nil {
		h.right.postOrder()
	}

	//输出根节点
	fmt.Printf("no=%d, name=%s\n", h.no, h.name)
}

//preSelect 根据编号前序查找
func (h *heroNode) preSelect(no int) *heroNode {
	if h.no == no {
		return h
	}

	var temp *heroNode

	//递归左查找
	if h.left != nil {
		temp = h.left.preSelect(no)
	}

	if temp != nil {
		return temp
	}

	//递归右查找
	if h.right != nil {
		temp = h.right.preSelect(no)
	}
	return temp
}
