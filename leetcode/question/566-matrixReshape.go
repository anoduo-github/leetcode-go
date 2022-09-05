package question

//566. 重塑矩阵
func matrixReshape(mat [][]int, r int, c int) [][]int {
	a := len(mat)
	b := len(mat[0])
	if (a * b) != (r * c) {
		return mat
	}
	//创建结果
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	//转为一维数组
	temp := make([]int, 0)
	for i := 0; i < a; i++ {
		temp = append(temp, mat[i]...)
	}
	//重新赋值
	x, y := 0, 0
	for k, v := range temp {
		res[x][y] = v
		if (k+1)%c == 0 {
			x++
			y = 0
		} else {
			y++
		}
	}
	return res
}
