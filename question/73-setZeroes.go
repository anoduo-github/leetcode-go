package question

//73. 矩阵置零
func setZeroes(matrix [][]int) {
	//行数
	row := len(matrix)
	//列数
	col := len(matrix[0])
	//定义两个变量记录首行和首列的情况
	row0 := false
	col0 := false
	//遍历首行首列
	for _, v := range matrix[0] {
		if v == 0 {
			row0 = true
			break
		}
	}
	for _, v := range matrix {
		if v[0] == 0 {
			col0 = true
			break
		}
	}
	//之后将首行首列用作记录其余行列中0的位置
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}
	//遍历首行首列，将0所处的行和列置0
	for i, v := range matrix[0] {
		if v == 0 {
			for j := 1; j < row; j++ {
				matrix[j][i] = 0
			}
		}
	}
	for i := 1; i < row; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < col; j++ {
				matrix[i][j] = 0
			}
		}
	}
	//判断是否将首行首列置0
	if row0 {
		for i := 0; i < col; i++ {
			matrix[0][i] = 0
		}
	}
	if col0 {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
