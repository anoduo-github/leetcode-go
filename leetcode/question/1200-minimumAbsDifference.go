package question

import (
	"math"
	sor "sort"
)

//1200. 最小绝对差
func minimumAbsDifference(arr []int) [][]int {
	res := make([][]int, 0)
	//先排序
	sor.Ints(arr)
	//遍历
	max := math.MaxInt32
	for i := 0; i < len(arr)-2; i++ {
		dest := arr[i+1] - arr[i]
		if dest < max {
			max = dest
			//表示这个区间比之前的值都小，重置结果
			res = [][]int{{arr[i], arr[i+1]}}
		} else if dest == max {
			//表示一致，往结果添加
			res = append(res, []int{arr[i], arr[i+1]})
		}
	}
	return res
}
