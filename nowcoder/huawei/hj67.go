package huawei

import (
	"fmt"
)

//HJ67 24点游戏算法
func HJ67(nums []int) {
	for i := 0; i < len(nums); i++ {
		temp := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if j != i {
				temp = append(temp, nums[j])
			}
		}
		if dfs(nums[i], temp) {
			fmt.Println(true)
			return
		}
	}
	fmt.Println(false)
}

func dfs(n int, nums []int) bool {
	if len(nums) == 0 {
		return n == 24
	}
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		temp := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if j != i {
				temp = append(temp, nums[j])
			}
		}
		if dfs(n*a, temp) {
			return true
		}
		if nums[i] != 0 && dfs(n/a, temp) {
			return true
		}
		if dfs(n+a, temp) {
			return true
		}
		if dfs(n-a, temp) {
			return true
		}
	}
	return false
}
