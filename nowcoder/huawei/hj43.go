package huawei

import (
	"fmt"
)

var (
	flag = false
)

//HJ43 迷宫问题
func HJ43(nums [][]int) {
	//创建一个临时数组，判断对应位置的值是否访问过
	isVisited := make([][]int, len(nums))
	copy(isVisited, nums)
	res := make([][]int, 0)
	walk(0, 0, nums, isVisited, res)
}

func walk(i, j int, nums, isVisited, res [][]int) {
	if isVisited[i][j] == -1 || flag {
		return
	}
	if i == len(nums)-1 && j == len(nums[0])-1 {
		//表示是最后一个点
		res = append(res, []int{i, j})
		for _, v := range res {
			fmt.Printf("(%d,%d)\n", v[0], v[1])
		}
		flag = true
		return
	}
	if nums[i][j] == 1 {
		return
	}
	res = append(res, []int{i, j})
	isVisited[i][j] = -1
	//向上移动
	if i > 0 {
		temp := make([][]int, len(isVisited))
		copy(temp, isVisited)
		ans := make([][]int, len(res))
		copy(ans, res)
		walk(i-1, j, nums, temp, ans)
	}
	//向下移动
	if i+1 < len(nums) {
		temp := make([][]int, len(isVisited))
		copy(temp, isVisited)
		ans := make([][]int, len(res))
		copy(ans, res)
		walk(i+1, j, nums, temp, ans)
	}
	//向左移动
	if j > 0 {
		temp := make([][]int, len(isVisited))
		copy(temp, isVisited)
		ans := make([][]int, len(res))
		copy(ans, res)
		walk(i, j-1, nums, temp, ans)
	}
	//向右移动
	if j+1 < len(nums[0]) {
		temp := make([][]int, len(isVisited))
		copy(temp, isVisited)
		ans := make([][]int, len(res))
		copy(ans, res)
		walk(i, j+1, nums, temp, ans)
	}
}
