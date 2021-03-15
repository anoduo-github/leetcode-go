package question

//54. 螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	//想法：一层层输出，按上，右，下，左的顺序

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	//定义结果集
	res := make([]int, 0)
	//行
	row := len(matrix)
	//列
	col := len(matrix[0])
	//定义左右和上下区间范围
	left, right, top, low := 0, col-1, 0, row-1
	for left <= right && top <= low {
		//上
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		//右
		for j := top + 1; j <= low; j++ {
			res = append(res, matrix[j][right])
		}
		//防止重复
		if left < right && top < low {
			//下
			for k := right - 1; k >= left; k-- {
				res = append(res, matrix[low][k])
			}
			//左
			for l := low - 1; l >= top+1; l-- {
				res = append(res, matrix[l][left])
			}
		}
		//进行下一层
		left++
		right--
		top++
		low--
	}
	return res
}
