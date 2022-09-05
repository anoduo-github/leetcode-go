package question

//155. 最小栈

type MinStack struct {
	data []int
	min  []int //存最小值的栈
}

/** initialize your data structure here. */
func Constructor33() MinStack {
	return MinStack{
		data: make([]int, 0),
		min:  make([]int, 0),
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.min) == 0 {
		this.min = append(this.min, x)
	} else {
		if this.min[len(this.min)-1] > x {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, this.min[len(this.min)-1])
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		this.data = this.data[0 : len(this.data)-1]
		this.min = this.min[0 : len(this.min)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) == 0 {
		return 0
	}
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}
