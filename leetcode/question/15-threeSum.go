package question

import (
	"fmt"
	sort2 "sort"
)

//15. 三数之和
func threeSum(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	//排序
	sort2.Ints(nums) //方便去重
	//先从左往右选第一个目标值a
	for a := 0; a < n && nums[a] <= 0; a++ {
		if a > 0 && nums[a] == nums[a-1] { //表示左右两个数相同，跳过
			continue
		}
		c := n - 1
		//b为a右边第一数，循环递增
		for b := a + 1; b < n; b++ {
			if b > a+1 && nums[b] == nums[b-1] { //表示左右两个数相同，跳过
				continue
			}
			//c为右边第一数，循环递减，且大于等于b
			for ; c > b; c-- {
				if nums[a]+nums[b]+nums[c] == 0 {
					//符合条件
					res = append(res, []int{nums[a], nums[b], nums[c]})
					break
				} else if nums[a]+nums[b]+nums[c] < 0 {
					//表示b,c的值加起来比a的绝对值小，又因为c已是最大值，所以break循环，b递增，增大b的值，重新比较
					break
				} else {
					//表示b,c的值加起来大于a的绝对值,所以继续循环递减c,使b+c的值逐渐变小
					continue
				}
			}
		}
	}
	return res
}

func Example15() {
	test := []int{-2, 0, 1, 1, 2}
	res := threeSum(test)
	fmt.Println(res)
}
