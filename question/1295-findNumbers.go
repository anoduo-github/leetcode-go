package question

import "strconv"

//1295. 统计位数为偶数的数字
func findNumbers(nums []int) int {
	count := 0
	for _, v := range nums {
		vt := strconv.Itoa(v)
		if len(vt)%2 == 0 {
			count++
		}
	}
	return count
}
