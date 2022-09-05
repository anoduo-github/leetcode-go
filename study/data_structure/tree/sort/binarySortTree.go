package sort

import "fmt"

//node 节点
type node struct {
	val   int   //权值
	left  *node //左子节点
	right *node //右子节点
}

//add 添加节点
func (n *node) add(e *node) {
	//添加的节点小于当前节点,应放在当前节点左边
	if e.val < n.val {
		//当前节点左边为空，直接放置
		if n.left == nil {
			n.left = e
		} else {
			//当前节点的左子节点已经存在了，递归添加
			n.left.add(e)
		}
	} else {
		//添加的节点大于等于当前节点,应放在当前节点右边
		if n.right == nil {
			n.right = e
		} else {
			n.right.add(e)
		}
	}
}

//findParent 根据val寻找目标节点和它的父节点
func (n *node) findParent(target int) (*node, *node, int) {
	//最后返回的int,0表示当前节点，1表示右子节点，-1表示左子节点
	if n.val > target {
		if n.left != nil {
			if n.left.val == target {
				return n.left, n, -1
			} else {
				return n.left.findParent(target)
			}
		} else {
			return nil, nil, 0
		}
	} else if n.val < target {
		if n.right != nil {
			if n.right.val == target {
				return n.right, n, 1
			} else {
				return n.right.findParent(target)
			}
		} else {
			return nil, nil, 0
		}
	} else {
		return n, nil, 0
	}
}

//midOrder 中序遍历
func (n *node) midOrder() {
	if n.left != nil {
		n.left.midOrder()
	}
	fmt.Printf("%d -> ", n.val)
	if n.right != nil {
		n.right.midOrder()
	}
}

//binarySortTree 二叉排序树
type binarySortTree struct {
	root *node
}

//Add 添加节点
func (b *binarySortTree) Add(n *node) {
	if b.root == nil {
		b.root = n
	} else {
		b.root.add(n)
	}
}

//MidOrder 中序遍历
func (b *binarySortTree) MidOrder() {
	if b.root != nil {
		b.root.midOrder()
		fmt.Println()
	}
}

//RemoveNode 删除节点
func (b *binarySortTree) RemoveNode(target int) {
	if b.root == nil {
		return
	}
	//先找到目标节点和它的父节点
	targetNode, parentNode, tag := b.root.findParent(target)

	//表示没有此节点，直接返回
	if targetNode == nil {
		fmt.Printf("没有val=%d的节点\n", target)
		return
	}

	//表示目标节点就是根节点
	if targetNode != nil && parentNode == nil {
		fmt.Printf("val=%d的节点是根节点,已删除\n", target)
		b.root = nil
		return
	}
	//接下来就是目标节点和父节点都存在的情况了

	//如果是叶子节点
	if targetNode.left == nil && targetNode.right == nil {
		if tag == -1 {
			fmt.Printf("val=%d的节点是叶子节点,是val=%d的节点的左子节点,已删除\n", target, parentNode.val)
			parentNode.left = nil
		} else if tag == 1 {
			fmt.Printf("val=%d的节点是叶子节点,是val=%d的节点的右子节点,已删除\n", target, parentNode.val)
			parentNode.right = nil
		}
		return
	}

	//如果目标节点有一个子节点
	//存在左子节点
	if targetNode.left != nil && targetNode.right == nil {
		if tag == -1 {
			fmt.Printf("val=%d的节点存在val=%d的左子节点,而且是val=%d的节点的左子节点,已删除\n", target, targetNode.left.val, parentNode.val)
			parentNode.left = targetNode.left
		} else if tag == 1 {
			fmt.Printf("val=%d的节点存在val=%d的左子节点,而且是val=%d的节点的右子节点,已删除\n", target, targetNode.left.val, parentNode.val)
			parentNode.right = targetNode.left
		}
	} else if targetNode.left == nil && targetNode.right != nil {
		if tag == -1 {
			fmt.Printf("val=%d的节点存在val=%d的右子节点,而且是val=%d的节点的左子节点,已删除\n", target, targetNode.right.val, parentNode.val)
			parentNode.left = targetNode.right
		} else if tag == 1 {
			fmt.Printf("val=%d的节点存在val=%d的右子节点,而且是val=%d的节点的右子节点,已删除\n", target, targetNode.right.val, parentNode.val)
			parentNode.right = targetNode.right
		}
	} else if targetNode.left != nil && targetNode.right != nil {
		//从右子节点遍历找到一个最小的节点，删除这个节点，将值赋给目标节点，重构排序树
		tempParent := targetNode
		temp := targetNode.right
		for temp.left != nil {
			tempParent = temp
			temp = temp.left
		}
		//删除
		if tempParent.left.val == temp.val {
			tempParent.left = nil
		} else {
			tempParent.right = nil
		}
		targetNode.val = temp.val
		fmt.Printf("val=%d的节点存在左、右子节点,已删除\n", target)
	}
}

//TestBinarySortTree 测试二叉排序树
func TestBinarySortTree() {
	bTree := &binarySortTree{}
	nums := []int{7, 3, 10, 12, 5, 1, 9, 11}
	for _, v := range nums {
		n := &node{
			val: v,
		}
		bTree.Add(n)
	}

	fmt.Println("*********删除节点**********")

	bTree.RemoveNode(10)

	fmt.Println("***********中序遍历********")
	bTree.MidOrder()
}
