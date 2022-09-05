package question

import sss "sort"

//1331. 数组序号转换
func arrayRankTransform(arr []int) []int {
	temp := make([]int, len(arr))
	copy(temp, arr)
	//排序
	sss.Ints(temp)
	//map存序号和值
	m := make(map[int]int, 0)
	i := 1
	for _, v := range temp {
		if _, ok := m[v]; !ok {
			m[v] = i
			i++
		}
	}
	//得结果
	res := make([]int, 0)
	for _, v := range arr {
		res = append(res, m[v])
	}
	return res
}
