package huffman

//huffmanTree 赫夫曼树
type huffmanTree struct {
	root *node
}

//NewHuffmanTree 新建一个huffmanTree
func NewHuffmanTree(list []*node) *huffmanTree {
	for len(list) > 1 {
		//排序
		nodeSort(list)
		//取出最小的两个
		one := list[0]
		two := list[1]
		//缩小列表
		list = list[2:]
		//新建一个新的节点
		newNode := new(node)
		newNode.left = one
		newNode.right = two
		newNode.val = one.val + two.val
		//放入节点
		list = append(list, newNode)
	}
	return &huffmanTree{
		root: list[0],
	}
}

func TestHuffmanTree() {
	node1 := &node{
		val: 1,
	}
	node2 := &node{
		val: 2,
	}
	node3 := &node{
		val: 7,
	}
	node4 := &node{
		val: 10,
	}
	node5 := &node{
		val: 5,
	}
	list := []*node{node1, node2, node3, node4, node5}
	ht := NewHuffmanTree(list)
	ht.root.preOrder()
}
