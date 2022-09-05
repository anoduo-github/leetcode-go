package huffman

import (
	"fmt"
)

//node 树节点
type node struct {
	val   int   //权值
	left  *node //左子节点
	right *node //右子节点
}

//NewNode 创建节点
func NewNode() *node {
	return &node{
		val: 0,
	}
}

//preOrder 前序遍历
func (n *node) preOrder() {
	//打印当前节点
	fmt.Printf("%d -> ", n.val)
	//递归查找左
	if n.left != nil {
		n.left.preOrder()
	}
	//递归查找右
	if n.right != nil {
		n.right.preOrder()
	}
}

//nodeSort 节点排序（从小到大）
func nodeSort(list []*node) {
	//采用希尔排序
	length := len(list)
	for d := length / 2; d > 0; d = d / 2 {
		for i := d; i < length; i = i + d {
			temp := list[i]
			j := i - d
			if list[j].val > temp.val {
				for j >= 0 && list[j].val > temp.val {
					list[j+d] = list[j]
					j = j - d
				}
				list[j+d] = temp
			}
		}
	}
}
