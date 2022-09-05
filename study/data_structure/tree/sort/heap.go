package sort

//Heap 堆
type heap struct {
	size  int   //堆大小
	array []int //用数组表示树
}

//NewHeap 新建一个堆
func NewHeap(list []int) *heap {
	return &heap{
		size:  0,
		array: list,
	}
}

//Push 添加一个
func (h *heap) Push(target int) {
	//如果大小为空的话
	if h.size == 0 {
		h.array[0] = target
		h.size++
		return
	}

	//要插入的下标为
	i := h.size

	for i > 0 {
		//定义父节点
		parent := (i - 1) / 2

		//如果插入的值小于父节点
		if target <= h.array[parent] {
			//中断循环
			break
		}

		//否则的话将父节点的值下移
		h.array[i] = h.array[parent]

		//移动i,继续比较,直到target小于它所在位置的父节点
		i = parent
	}

	//赋值，相当于此时的i就是target该插入的位置
	h.array[i] = target
	h.size++
}

//Pop 弹出最大元素,这个方法只缩短size的大小，数组大小不改变，将被弹的值放在数组最后面，通过size直接访问不到
func (h *heap) Pop() int {
	if h.size == 0 {
		return -1
	}
	//取出根节点
	root := h.array[0]

	//因为根节点弹出，将最后一个节点放入根节点
	h.size--
	temp := h.array[h.size]
	h.array[h.size] = root

	//比较，将大的值往上移，小的值下沉
	i := 0
	for {
		a := 2*i + 1
		b := 2*i + 2

		if a >= h.size {
			//超标
			break
		}

		//找到最大子节点的下标
		if b < h.size && h.array[b] > h.array[a] {
			a = b
		}

		//如果父节点的值大于或等于左右子节点的值，中止循环
		if temp >= h.array[a] {
			break
		}

		//与父节点交换
		h.array[i] = h.array[a]

		i = a
	}

	h.array[i] = temp

	return root
}
