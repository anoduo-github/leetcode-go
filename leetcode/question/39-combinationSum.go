package question

//39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, 0)
	var dfs func(index, target int)
	dfs = func(index, target int) {
		if index == len(candidates) {
			//超出数组范围了
			return
		}
		if target == 0 {
			res = append(res, append([]int(nil), temp...))
			return
		}
		//跳过当前值
		dfs(index+1, target)
		//选择当前值
		if target-candidates[index] >= 0 {
			//不满足上面条件表示选择当前值达不到和相等的要求
			temp = append(temp, candidates[index])
			dfs(index, target-candidates[index])
			temp = temp[0 : len(temp)-1] //这一步相当于回溯，返回之前的状态
		}
	}
	dfs(0, target)
	return res
}
