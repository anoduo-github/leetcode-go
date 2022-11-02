package question

//46. 全排列
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	visited := make([]bool, len(nums))
	var dfs func(temp []int)
	dfs = func(temp []int) {
		if len(temp) == len(nums) {
			temp2 := make([]int, len(temp))
			copy(temp2, temp)
			res = append(res, temp2)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				temp = append(temp, nums[i])
				visited[i] = true
				dfs(temp)
				//回溯
				visited[i] = false
				temp = temp[0 : len(temp)-1]
			}
		}
	}
	dfs(make([]int, 0))
	return res
}
