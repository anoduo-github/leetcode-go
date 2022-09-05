package question

import "math"

// 1779. 找到最近的有相同 X 或 Y 坐标的点
func nearestValidPoint(x int, y int, points [][]int) int {
	index := -1
	min := math.MaxInt
	//遍历
	for i := 0; i < len(points); i++ {
		if x == points[i][0] || y == points[i][1] {
			d := distance(x, y, points[i][0], points[i][1])
			if min > d {
				min = d
				index = i
			} else if min == d && index > i {
				index = i
			}
		}
	}
	return index
}

//distance 曼哈顿距离
func distance(x, y, i, j int) int {
	return int(math.Abs(float64(x)-float64(i)) + math.Abs(float64(y)-float64(j)))
}
