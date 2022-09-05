package interview

import "testing"

func TestSearch(t *testing.T) {
	matrix := make([][]int, 5)
	matrix[0] = []int{1, 4, 7, 11, 15}
	matrix[1] = []int{2, 5, 8, 12, 19}
	matrix[2] = []int{3, 6, 9, 16, 22}
	matrix[3] = []int{10, 13, 14, 17, 24}
	matrix[4] = []int{18, 21, 23, 26, 30}
	res := searchMatrix(matrix, 5)
	t.Error(res)
}
