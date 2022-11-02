package huawei

import (
	"fmt"
	"strconv"
)

//HJ57 高精度整数加法
func HJ57() {
	var x string
	fmt.Scanln(&x)
	var y string
	fmt.Scanln(&y)

	lx := len(x)
	ly := len(y)
	var max int
	if lx > ly {
		max = lx + 1
	} else {
		max = ly + 1
	}
	res := make([]int, max)

	i, j, z := lx-1, ly-1, max-1
	temp := 0
	for i >= 0 || j >= 0 {
		a := 0
		if i >= 0 {
			a, _ = strconv.Atoi(string(x[i]))
		}
		b := 0
		if j >= 0 {
			b, _ = strconv.Atoi(string(y[j]))
		}
		sum := a + b + temp
		temp = sum / 10
		res[z] = sum % 10
		z--
		i--
		j--
	}
	if temp > 0 {
		res[z] = temp
	}
	for i, v := range res {
		if i > 0 {
			fmt.Print(v)
		} else {
			if v > 0 {
				fmt.Print(v)
			}
		}
	}
}
