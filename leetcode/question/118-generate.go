package question

//118. 杨辉三角
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	if numRows == 1 {
		res[0] = []int{1}
		return res
	} else if numRows == 2 {
		res[0] = []int{1}
		res[1] = []int{1, 1}
		return res
	} else {
		res[0] = []int{1}
		res[1] = []int{1, 1}
		for i := 2; i < numRows; i++ {
			res[i] = make([]int, i+1)
			for j := 0; j < i+1; j++ {
				if j == 0 || j == i {
					res[i][j] = 1
				} else {
					res[i][j] = res[i-1][j-1] + res[i-1][j]
				}
			}
		}
	}
	return res
}
