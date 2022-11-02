package question

import (
	"sort"
)

//1619. 删除某些元素后的数组均值
func trimMean(arr []int) float64 {
	sort.Ints(arr)
	length := len(arr)
	sum := 0
	for _, v := range arr[length/20 : length*19/20] {
		sum += v
	}
	return float64(sum) / (float64(length) * 0.9)
}
