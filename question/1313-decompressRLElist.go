package question

//1313. 解压缩编码列表
func decompressRLElist(nums []int) []int {
	res := make([]int, 0)
	i := 0
	for i < len(nums) {
		for j := 0; j < nums[i]; j++ {
			res = append(res, nums[i+1])
		}
		i = i + 2
	}
	return res
}
