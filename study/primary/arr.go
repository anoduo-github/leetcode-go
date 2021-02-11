package primary

//LC 删除排序数组中的重复项
func removeDuplicates(nums []int) int {
	//双指针
	length := len(nums)
	if length <= 1 {
		return length
	}
	i := 0
	for j := 1; j < length; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

//LC 买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	//采用双指针,因为比较大小即可，左边比右边小就是有利润
	//先获取长度
	l := len(prices)
	//定义利润
	sum := 0
	//定义指针
	i := 0
	j := 1
	for j < l {
		//当左边比右边大，指针前移
		if prices[i] >= prices[j] {
			i = j
			j++
			continue
		}
		//当左边比右边小，继续比较右边与右边的右边，直到找到右边连续增大的最后一个
		if j < l-1 && prices[j] < prices[j+1] {
			j++
		} else {
			sum = sum + prices[j] - prices[i]
			i = j + 1
			j += 2
		}
	}
	return sum
}

//LC 有效的数独
func isValidSudoku(board [][]byte) bool {
	//行
	row := make([]map[int]int, 9)
	//列
	col := make([]map[int]int, 9)
	//3x3区域
	area := make([]map[int]int, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - '0')
			//初始化map,因为在循环体里，加判断防止重复初始化
			if row[i] == nil {
				row[i] = make(map[int]int)
			}
			//判断行是否重复
			if _, ok := row[i][num]; ok {
				return false
			}
			row[i][num] = 1
			//初始化map
			if col[j] == nil {
				col[j] = make(map[int]int)
			}
			//判断列是否重复
			if _, ok := col[j][num]; ok {
				return false
			}
			col[j][num] = 1
			//初始化map
			if area[j/3+(i/3)*3] == nil {
				area[j/3+(i/3)*3] = make(map[int]int)
			}
			//判断区域是否存在重复,j/3是表示将九个元素分为3组，每组存放在一个area[index]中
			if _, ok := area[j/3+(i/3)*3][num]; ok {
				return false
			}
			area[j/3+(i/3)*3][num] = 1
		}
	}
	return true
}

//LC 旋转图像
func rotate(matrix [][]int) {
	//获取长度
	n := len(matrix)
	//从外向内旋转，找规律
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-1-j]
		}
	}
}
