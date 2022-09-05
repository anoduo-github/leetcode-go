package avl

import "fmt"

//avlTree AVL树
type avlTree struct {
	root *avlTreeNode //根节点
}

//NewAVLTree() 创建一个avl树
func NewAVLTree() *avlTree {
	return new(avlTree)
}

//Add 添加一个节点
func (a *avlTree) Add(val int) {
	a.root = a.root.add(val)
}

//Delete 删除一个节点
func (a *avlTree) Delete(target int) {
	if a == nil || a.root == nil {
		return
	}
	a.root = a.root.delete(target)
}

//FindMinValue 查找最小值
func (a *avlTree) FindMinValue() int {
	if a.root == nil {
		panic("AVLTree is nil")
	}
	return a.root.findMinValue()
}

//FindMaxValue 查找最大值
func (a *avlTree) FindMaxValue() int {
	if a.root == nil {
		panic("AVLTree is nil")
	}
	return a.root.findMaxValue()
}

//Find 查找指定值的节点
func (a *avlTree) Find(target int) *avlTreeNode {
	if a.root == nil {
		panic("AVLTree is nil")
	}
	return a.root.find(target)
}

//MidOrder 中序遍历
func (a *avlTree) MidOrder() {
	if a.root == nil {
		return
	}
	a.root.midOrder()
	fmt.Println()
}

//TestAVLTree 测试avl树
func TestAVLTree() {
	nums := []int{3, 5, 7, 5, 2, 1, 6, 9, 10, 8, 45, 23}
	at := NewAVLTree()
	for _, v := range nums {
		at.Add(v)
	}
	at.MidOrder()

	fmt.Println("执行删除")

	at.Delete(9)
	at.MidOrder()
}
