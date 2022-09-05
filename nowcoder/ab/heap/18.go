package heap

import (
	"errors"
)

//AB18 【模板】堆
type Heap struct {
	array []int //存数据
	size  int   //存大小
}

func NewHeap() *Heap {
	return &Heap{
		array: make([]int, 0),
		size:  0,
	}
}

//adjustHeap 调整堆
func (h *Heap) adjustHeap(i int) {
	if h.size == 0 {
		return
	}
	temp := h.array[i]
	//比较当前节点的左子节点和右子节点
	for k := 2*i + 1; k < h.size; k = 2*k + 1 {
		if k+1 < h.size && h.array[k+1] > h.array[k] {
			//表示右子节点比左子节点大，将下标移到右子节点
			k++
		}
		if h.array[k] > temp {
			//表示当前节点小于自己的子节点，根据大顶堆，根节点要是最大的，所以要互换
			h.array[i] = h.array[k]
			//移到重新赋值的节点上
			i = k
		} else {
			//表示符合大顶堆
			break
		}
	}
	//别忘了这一步，当发生改变时给当前节点重新赋值
	h.array[i] = temp
}

func (h *Heap) Push(x int) {
	//将新元素放入到末尾
	h.array = append(h.array, x)
	h.size++
	//调整最后一个节点所在的子树
	for i := h.size/2 - 1; i >= 0; i = (i - 1) / 2 {
		h.adjustHeap(i)
		//由于(0-1)/2=0,防止i=0死循环
		if i == 0 {
			break
		}
	}
}

func (h *Heap) Pop() (int, error) {
	if h.size > 0 {
		temp := h.array[0]
		h.array[0] = h.array[h.size-1]
		h.array = h.array[0 : h.size-1]
		h.size--
		//从第一个开始调整
		h.adjustHeap(0)
		return temp, nil
	} else {
		return 0, errors.New("empty")
	}
}

func (h *Heap) Top() (int, error) {
	if h.size > 0 {
		return h.array[0], nil
	} else {
		return 0, errors.New("empty")
	}
}
