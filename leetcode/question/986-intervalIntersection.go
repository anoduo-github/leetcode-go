package question

//986. 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0)
	f := len(firstList)
	s := len(secondList)
	if f == 0 || s == 0 {
		return res
	}
	i, j := 0, 0
	for i < f && j < s {
		min := 0
		max := 0
		if firstList[i][0] <= secondList[j][0] {
			min = secondList[j][0]
		} else {
			min = firstList[i][0]
		}
		if firstList[i][1] <= secondList[j][1] {
			max = firstList[i][1]
			i++
		} else {
			max = secondList[j][1]
			j++
		}
		if min <= max {
			res = append(res, []int{min, max})
		}
	}
	return res
}
