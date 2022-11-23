package huawei

import "fmt"

//HJ69 矩阵乘法
func HJ69() {
	var x, y, z int
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	fmt.Scanln(&z)
	A := make([][]int, x)
	B := make([][]int, y)
	C := make([][]int, x)
	for i := 0; i < x; i++ {
		C[i] = make([]int, z)
		row := make([]int, y)
		for j := 0; j < y; j++ {
			var num int
			fmt.Scan(&num)
			row[j] = num
		}
		A[i] = row
	}
	for i := 0; i < y; i++ {
		row := make([]int, z)
		for j := 0; j < z; j++ {
			var num int
			fmt.Scan(&num)
			row[j] = num
		}
		B[i] = row
	}

	for i := 0; i < x; i++ {
		for j := 0; j < z; j++ {
			temp := 0
			for k := 0; k < y; k++ {
				temp += A[i][k] * B[k][j]
			}
			C[i][j] = temp
		}
	}

	for _, v := range C {
		for i := 0; i < z; i++ {
			fmt.Print(v[i])
			if i < z-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
