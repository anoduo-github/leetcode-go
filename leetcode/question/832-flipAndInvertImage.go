package question

//flipAndInvertImage 翻转图像
func flipAndInvertImage(A [][]int) [][]int {
	for i := 0; i < len(A); i++ {
		reverseOrder(A[i])
		reversal(A[i])
	}
	return A
}

//reverseOrder 逆序
func reverseOrder(B []int) {
	//双指针
	//左指针
	i := 0
	//右指针
	j := len(B) - 1
	for i < j {
		B[i], B[j] = B[j], B[i]
		i++
		j--
	}
}

//reversal 反转(0转1)
func reversal(C []int) {
	for i := 0; i < len(C); i++ {
		if C[i] == 0 {
			C[i] = 1
		} else {
			C[i] = 0
		}
	}
}
