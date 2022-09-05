package offer

//剑指 Offer 09. 用两个栈实现队列

type CQueue struct {
	in  []int //入栈
	out []int //出栈
}

func Constructor2() CQueue {
	return CQueue{
		in:  make([]int, 0),
		out: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.out) == 0 {
		//如果出栈为空
		if len(this.in) == 0 {
			//如果入栈也为空，则报错
			return -1
		}
		//否则将入栈元素加入出栈
		this.out = append(this.out, this.in...)
		this.in = make([]int, 0)
	}
	res := this.out[0]
	this.out = this.out[1:]
	return res
}
