package question

//448. 找到所有数组中消失的数字
func findDisappearedNumbers(nums []int) []int {
	res := make([]int, 0)
	temp := make([]int, len(nums)+1)
	for _, v := range nums {
		temp[v]++
	}
	for i, v := range temp {
		if v == 0 && i > 0 {
			res = append(res, i)
		}
	}
	return res
}
