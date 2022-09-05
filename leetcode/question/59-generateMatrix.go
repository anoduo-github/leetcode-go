package question

//59. 螺旋矩阵 II

func GenerateMatrix(n int) [][]int {
	//创建数据
	nums := make([]int, 0)
	for i := 1; i <= n*n; i++ {
		nums = append(nums, i)
	}
	//创建二维矩阵
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	//定义左右和上下区间范围
	left, right, top, low := 0, n-1, 0, n-1
	//定义数据组下标
	index := 0
	for left <= right && top <= low {
		//上
		for i := left; i <= right; i++ {
			matrix[top][i] = nums[index]
			index++
		}
		//右
		for j := top + 1; j <= low; j++ {
			matrix[j][right] = nums[index]
			index++
		}
		//防止重复
		if left < right && top < low {
			//下
			for k := right - 1; k >= left; k-- {
				matrix[low][k] = nums[index]
				index++
			}
			//左
			for l := low - 1; l >= top+1; l-- {
				matrix[l][left] = nums[index]
				index++
			}
		}
		//进行下一层
		left++
		right--
		top++
		low--
	}
	return matrix
}
