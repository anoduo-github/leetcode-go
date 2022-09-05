package question

type MyCircularQueue struct {
	data  []int //切片存数据
	size  int   //已用容量
	cap   int   //总容量
	index int   //队首下标
	last  int   //队尾下标
}

func Constructor22(k int) MyCircularQueue {
	return MyCircularQueue{
		data:  make([]int, k),
		size:  0,
		cap:   k,
		index: 0,
		last:  -1,
	}
}

//EnQueue 插入元素
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == this.cap {
		return false
	}
	//添加
	this.last = (this.last + 1) % this.cap
	this.data[this.last] = value
	this.size++
	return true
}

//DeQueue 删除元素
func (this *MyCircularQueue) DeQueue() bool {
	//删除队首，先进先出
	if this.size == 0 {
		return false
	}
	this.index = (this.index + 1) % this.cap
	this.size--
	return true
}

//Front 获取队首元素
func (this *MyCircularQueue) Front() int {
	if this.size == 0 {
		return -1
	}
	return this.data[this.index]
}

//Rear 获取队尾元素
func (this *MyCircularQueue) Rear() int {
	if this.size == 0 {
		return -1
	}
	return this.data[this.last]
}

//IsEmpty 是否为空
func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

//IsFull 是否已满
func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.cap
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
