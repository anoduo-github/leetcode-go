package question

import "math/rand"

//380. O(1) 时间插入、删除和获取随机元素

type RandomizedSet struct {
	nums []int       //存值
	m    map[int]int //存值对应的下标
}

func Constructor10() RandomizedSet {
	return RandomizedSet{
		nums: make([]int, 0),
		m:    make(map[int]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		//表示值已经存在
		return false
	}
	//先存值
	this.nums = append(this.nums, val)
	//存下标
	this.m[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	id, ok := this.m[val]
	if !ok {
		//表示值不存在
		return false
	}
	//将数组最后一位的值覆盖到被删除的值的位置，然后缩短数组长度，实现删除
	last := len(this.nums) - 1
	this.nums[id] = this.nums[last] //赋值
	this.m[this.nums[id]] = id      //将最后一位的值的下标改成被删除的值的下标，相当于替换了位置
	this.nums = this.nums[:last]    //数组移除最后一位
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}
