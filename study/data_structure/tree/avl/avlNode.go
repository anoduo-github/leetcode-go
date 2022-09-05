package avl

import "fmt"

//avlTreeNode avl树节点
type avlTreeNode struct {
	value  int          //值
	height int          //该节点作为树根节点时的树的高度
	left   *avlTreeNode //左子节点
	right  *avlTreeNode //右子节点
}

//updateHeight 更新节点树高度
func (node *avlTreeNode) updateHeight() {
	if node == nil {
		return
	}

	//初始
	leftHeight, rightHeight := 0, 0

	if node.left != nil {
		leftHeight = node.left.height
	}
	if node.right != nil {
		leftHeight = node.right.height
	}

	//比较，找出最大，然后加上自己
	if leftHeight > rightHeight {
		node.height = leftHeight + 1
	} else {
		node.height = rightHeight + 1
	}
}

//balanceFactor 计算平衡因子
func (node *avlTreeNode) balanceFactor() int {
	leftHeight, rightHeight := 0, 0
	if node.left != nil {
		leftHeight = node.left.height
	}
	if node.right != nil {
		rightHeight = node.right.height
	}
	return leftHeight - rightHeight
}

//rightRotation 右旋
func rightRotation(node *avlTreeNode) *avlTreeNode {
	//首先创建一个新节点存放当前节点的左字节点
	temp := node.left
	//然后让当前节点的左子字节指向当前节点的左子节点的右子节点
	node.left = temp.right
	//然后让新节点的右字节点指向当前节点.相当于把新节点作为根节点，原根节点变成了根节点的右子节点
	temp.right = node

	//当前节点和新节点发生了变化，重新计算高度,注意顺序
	node.updateHeight()
	temp.updateHeight()
	return temp
}

//leftRotation 左旋
func leftRotation(node *avlTreeNode) *avlTreeNode {
	//首先创建一个新的节点存放当前节点的右子节点
	temp := node.right
	//然后让当前节点的右子节点指向当前节点的右子节点的左子节点
	node.right = temp.left
	//然后让新节点的左子节点指向当前节点
	temp.left = node

	node.updateHeight()
	temp.updateHeight()

	return temp
}

//leftRightRotation 先左旋在右旋
func leftRightRotation(node *avlTreeNode) *avlTreeNode {
	//将当前节点的左子节点指向左旋后的节点
	node.left = leftRotation(node.left)
	//在整体右旋
	return rightRotation(node)
}

//rightLeftRotation 先右旋再左旋
func rightLeftRotation(node *avlTreeNode) *avlTreeNode {
	//将当前节点的右子节点指向右旋后的节点
	node.right = rightRotation(node.right)
	//在整体左旋
	return leftRotation(node)
}

//add 添加一个节点
func (node *avlTreeNode) add(val int) *avlTreeNode {
	//为空创建
	if node == nil {
		return &avlTreeNode{
			value:  val,
			height: 1,
		}
	}
	//相等跳过
	if node.value == val {
		return node
	}

	var temp *avlTreeNode
	//判断往左加还是往右加
	if val < node.value {
		//表示节点要添加到左子节点
		node.left = node.left.add(val)
		//计算平衡因子
		factor := node.balanceFactor() //左减右，此时为正
		//因为avl树左右子节点高度差不能大于1，因此如果左边已经比右边多1，那么再插入一个，就是2了，需要旋转
		//不存在大于2的情况，因为avl的定义就是不能超过1，所以超过1就会旋转，然后重新满足1的需求
		if factor == 2 {
			//表示要旋转
			if val < node.left.value {
				//直接右旋
				temp = rightRotation(node)
			} else {
				temp = leftRightRotation(node)
			}
		}
	} else {
		//表示节点要添加到右子节点
		node.right = node.right.add(val)
		factor := node.balanceFactor() //左减右，此时为负
		if factor == -2 {
			if val > node.right.value {
				//直接左旋
				temp = leftRotation(node)
			} else {
				temp = rightLeftRotation(node)
			}
		}
	}

	if temp == nil {
		//表示未发生旋转
		node.updateHeight()
		return node
	} else {
		temp.updateHeight()
		return temp
	}
}

//delete 删除指定节点
func (node *avlTreeNode) delete(target int) *avlTreeNode {
	if node == nil {
		return nil
	}
	if target < node.value {
		//表示值小于当前节点，从左子树删除
		node.left = node.left.delete(target)
		//更新节点高度
		node.left.updateHeight()
	} else if target > node.value {
		//表示值大于当前节点，从右子树删除
		node.right = node.right.delete(target)
		//更新节点高度
		node.right.updateHeight()
	} else {
		//表示找到该节点了
		//如果该节点没有左右子节点,直接返回nil
		if node.left == nil && node.right == nil {
			return nil
		}
		//如果该节点有左右子节点
		//递归里包含了重新平衡的操作，因此直接返回node
		if node.left != nil && node.right != nil {
			if node.left.height > node.right.height {
				//如果左子树高度大于右子树，找出左子树中最大的节点
				maxNode := node.left
				for maxNode.right != nil {
					maxNode = maxNode.right
				}
				//将当前节点的值替换为最大值节点
				node.value = maxNode.value
				//相当于当前这个要被删除的节点，与最大值节点发生替换
				//保证当前节点不存在后，结构符合规范，左小右大
				//删除这个最大值节点
				node.left = node.left.delete(maxNode.value)
				//更新树高度
				node.left.updateHeight()
			} else {
				//如果右子树高度大于左子树,找出右子树的最小的节点
				minNode := node.right
				for minNode.left != nil {
					minNode = minNode.left
				}
				node.value = minNode.value
				node.right = node.right.delete(minNode.value)
				node.right.updateHeight()
			}
		} else {
			//如果该节点只有左或右子节点
			//因为满足avl条件时，子节点必定是最下一层了，去除它之后，当前节点与最下一层依然只差1，符合条件，因此操作完毕直接返回node
			if node.left != nil {
				//存在左子节点
				//将值替换
				node.value = node.left.value
				//左子节点置空，因为根据avl,当只有一个子节点时，该节点肯定时叶子节点
				node.left = nil
				//设置高度,因为没有子节点了吗，这个高度是相对当前节点的嘛
				node.height = 1
			} else {
				//根据之前的判断已经剔除了都不为空、都存在的情况，所以这里直接else就行
				//存在右子节点
				node.value = node.right.value
				node.right = nil
				node.height = 1
			}
		}
		//执行完操作后返回新的当前节点
		return node
	}
	//再平衡
	var newRoot *avlTreeNode
	if node.balanceFactor() == 2 {
		//表示左子节点大于右子节点2，要重新平衡
		if node.left.balanceFactor() >= 0 {
			//表示左子节点的左子节点高度大于其右子节点,直接右旋
			newRoot = rightRotation(node)
		} else {
			//表示左子节点的左子节点高度小于其右子节点,先左旋再右旋
			newRoot = leftRightRotation(node)
		}
	} else if node.balanceFactor() == -2 {
		//表示右子节点大于左子节点2，要重新平衡
		if node.right.balanceFactor() <= 0 {
			//表示右子节点的左子节点高度小于其右子节点,直接左旋
			newRoot = leftRotation(node)
		} else {
			//表示右子节点的左子节点高度大于其右子节点,先右旋再左旋
			newRoot = rightLeftRotation(node)
		}
	}

	if newRoot != nil {
		newRoot.updateHeight()
		return newRoot
	}
	node.updateHeight()
	return node
}

//findMinValue 查找最小值
func (node *avlTreeNode) findMinValue() int {
	if node.left == nil {
		return node.value
	}
	return node.left.findMinValue()
}

//findMaxValue 查找最大值
func (node *avlTreeNode) findMaxValue() int {
	if node.right == nil {
		return node.value
	}
	return node.right.findMaxValue()
}

//find 查找指定节点
func (node *avlTreeNode) find(target int) *avlTreeNode {
	if node.value == target {
		return node
	} else if node.value > target {
		//往左找
		if node.left != nil {
			return node.left.find(target)
		}
		return nil
	} else {
		//往右找
		if node.right != nil {
			return node.right.find(target)
		}
		return nil
	}
}

//midOrder 中序遍历
func (node *avlTreeNode) midOrder() {
	if node.left != nil {
		node.left.midOrder()
	}
	fmt.Printf("%d -> ", node.value)
	if node.right != nil {
		node.right.midOrder()
	}
}
