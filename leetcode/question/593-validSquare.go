package question

//593. 有效的正方形
func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	//以p1为基准点
	l12 := (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
	l13 := (p1[0]-p3[0])*(p1[0]-p3[0]) + (p1[1]-p3[1])*(p1[1]-p3[1])
	l14 := (p1[0]-p4[0])*(p1[0]-p4[0]) + (p1[1]-p4[1])*(p1[1]-p4[1])
	if l13 == l14 && l12 > l13 {
		return l12 == (l13 + l14)
	} else if l12 == l14 && l13 > l12 {
		return l13 == (l12 + l14)
	} else if l12 == l13 && l14 > l12 {
		return l14 == (l12 + l13)
	}
	return false
}
