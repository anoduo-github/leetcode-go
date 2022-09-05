package clue

import "fmt"

//遍历线索化二叉树的实质还是在遍历二叉树，只是加了线索条件
//前序遍历线索化，还是先打印出父节点，在打印左子节点，然后打印右子节点或者后继节点
//中序遍历线索化，还是先打印前驱节点为空的头节点，然后打印后继节点和右子节点
//后序遍历线索化，还是先打印前驱节点为空的头节点，然后打印后继节点，然后父节点

//clueBinaryTree 线索化二叉树
type clueBinaryTree struct {
	root *heroNode //根节点
	pre  *heroNode //前驱节点,线索化和遍历需要用
}

//preClue 前序线索化
func (c *clueBinaryTree) preClue() {
	if c.root == nil {
		return
	}
	c.preclue(c.root)
}

//midClue 中序线索化
func (c *clueBinaryTree) midClue() {
	if c.root == nil {
		return
	}
	c.midclue(c.root)
}

//postClue 后序线索化
func (c *clueBinaryTree) postClue() {
	if c.root == nil {
		return
	}
	c.postclue(c.root)
}

//preclue 前序线索化
func (c *clueBinaryTree) preclue(node *heroNode) {
	//处理当前节点
	if node.left == nil { //表示左子节点为空
		node.left = c.pre //指向前驱节点，此时为空
		node.noLeft = 1   //标记节点状态
	}
	//处理当前节点的后继节点
	if c.pre != nil && c.pre.right == nil { //表示前驱节点不为空，即此时的前驱节点已经指向上次遍历到的节点
		c.pre.right = node //前驱节点的后继节点指向当前节点
		c.pre.noRight = 1  //标记状态
	}

	//前序遍历，当前节点就是上一个的前驱节点
	c.pre = node

	//递归左边
	if node.noLeft == 0 && node.left != nil { //0是找子节点，如果是1，就是前驱节点，那就是找之前已经处理过的节点了，这里显然不能找前驱节点
		c.preclue(node.left)
	}

	//递归右边
	if node.noRight == 0 && node.right != nil {
		c.preclue(node.right)
	}
}

//midclue 中序线索化，递归
func (c *clueBinaryTree) midclue(node *heroNode) {
	//递归左边
	if node.left != nil {
		c.midclue(node.left)
	}

	//处理当前节点
	if node.left == nil {
		node.left = c.pre
		node.noLeft = 1
	}
	//处理当前节点的后继节点
	if c.pre != nil && c.pre.right == nil {
		c.pre.right = node
		c.pre.noRight = 1
	}

	//中序遍历，当前节点就是上一个的前驱节点
	c.pre = node

	if node.right != nil {
		c.midclue(node.right)
	}
}

//postclue 后序线索化
func (c *clueBinaryTree) postclue(node *heroNode) {
	if node.left != nil {
		c.postclue(node.left)
	}
	if node.right != nil {
		c.postclue(node.right)
	}

	if node.left == nil {
		node.left = c.pre
		node.noLeft = 1
	}
	if c.pre != nil && c.pre.right == nil {
		c.pre.right = node
		c.pre.noRight = 1
	}
	c.pre = node
}

//printPreClue 打印前序遍历线索二叉树
func (c *clueBinaryTree) printPreClue() {
	temp := c.root

	for temp != nil {
		for temp.noLeft == 0 {
			fmt.Printf("no=%d,name=%s\n", temp.no, temp.name)
			temp = temp.left
		}
		fmt.Printf("no=%d,name=%s\n", temp.no, temp.name) //这是前驱节点
		temp = temp.right                                 //找后继节点
	}
}

//printMidClue 打印中序遍历线索二叉树
func (c *clueBinaryTree) printMidClue() {
	temp := c.root

	for temp != nil {
		//向左找头节点
		for temp.noLeft == 0 { //表示有左子节点
			temp = temp.left
		}

		fmt.Printf("no=%d,name=%s\n", temp.no, temp.name)

		for temp.noRight == 1 { //表示有后继节点
			temp = temp.right
			fmt.Printf("no=%d,name=%s\n", temp.no, temp.name)
		}

		temp = temp.right

	}
}

//printPostClue 打印后序遍历线索二叉树
func (c *clueBinaryTree) printPostClue() {
	temp := c.root

	//先找到头节点（也就是前驱节点为nil的）
	for temp != nil && temp.noLeft == 0 {
		temp = temp.left
	}

	for temp != nil {
		if temp.noRight == 1 { //表示有后继节点
			fmt.Printf("no=%d,name=%s\n", temp.no, temp.name)
			c.pre = temp //临时保存当前节点作为下一个节点的前驱节点，方便后续比较
			temp = temp.right
		} else {
			if temp.right == c.pre { //表示此节点为前驱节点的父节点,且前驱节点为右子节点
				fmt.Printf("no=%d,name=%s\n", temp.no, temp.name)
				/* if temp == c.root {
					return //执行完毕
				} */
				c.pre = temp
				temp = temp.parent //找父节点是因为此节点没有后继节点，只有右子节点，而右子节点已经遍历，别忘了这是后序遍历
			} else {
				temp = temp.right                     //找该节点的右子节点
				for temp != nil && temp.noLeft == 0 { //再次循环找该右子节点的左子节点
					temp = temp.left
				}
			}
		}
	}
}

//TestClue 测试线索化
func TestClue() {
	//构造树
	node1 := initBinaryTree()
	node2 := initBinaryTree()
	node3 := initBinaryTree()
	node4 := initBinaryTree()

	node4.delNode(3)
	node4.preOrder()
	t := node4.preSelect(14)
	fmt.Printf("t.name=%s\n", t.name)

	fmt.Println("*******前序遍历*******")
	node1.preOrder()
	fmt.Println("*******中序遍历*******")
	node1.midOrder()
	fmt.Println("*******后序遍历*******")
	node1.postOrder()

	fmt.Println("*******打印前序线索化二叉树*******")
	clueTree1 := new(clueBinaryTree)
	clueTree1.root = node1
	clueTree1.preClue()
	clueTree1.printPreClue()

	fmt.Println("*******打印中序线索化二叉树*******")
	clueTree2 := new(clueBinaryTree)
	clueTree2.root = node2
	clueTree2.midClue()
	clueTree2.printMidClue()

	fmt.Println("*******打印前序线索化二叉树*******")
	clueTree3 := new(clueBinaryTree)
	clueTree3.root = node3
	clueTree3.postClue()
	clueTree3.printPostClue()
}

//initBinaryTree 初始化一个二叉树
func initBinaryTree() *heroNode {
	//构造树
	root := &heroNode{
		no:   1,
		name: "吕布",
	}
	node1 := &heroNode{
		no:   3,
		name: "貂蝉",
	}
	node2 := &heroNode{
		no:   6,
		name: "曹操",
	}
	node3 := &heroNode{
		no:   8,
		name: "刘备",
	}
	node4 := &heroNode{
		no:   10,
		name: "关羽",
	}
	node5 := &heroNode{
		no:   14,
		name: "张飞",
	}

	root.left = node1
	root.right = node2
	node1.left = node3
	node1.right = node4
	node1.parent = root
	node2.left = node5
	node2.parent = root
	node3.parent = node1
	node4.parent = node1
	node5.parent = node2

	return root
}
