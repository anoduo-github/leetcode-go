package question

//1470. 重新排列数组
func shuffle(nums []int, n int) []int {
	//1.先将原数组分成两组
	numx := nums[:n]
	numy := nums[n:]
	//2.新建结果切片，遍历添加
	res := make([]int, 0)
	for i := 0; i < n; i++ {
		res = append(res, numx[i], numy[i])
	}
	return res
}
