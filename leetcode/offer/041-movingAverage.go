package offer

//剑指 Offer II 041. 滑动窗口的平均值

type MovingAverage struct {
	nums []int
	size int
	ans  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		nums: make([]int, 0),
		size: size,
		ans:  0,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.nums) == this.size {
		this.ans -= this.nums[0]
		this.nums = this.nums[1:]

	}
	this.nums = append(this.nums, val)
	this.ans += val

	return float64(this.ans) / float64(len(this.nums))
}
