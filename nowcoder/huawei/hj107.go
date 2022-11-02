package huawei

import (
	"fmt"
	"math"
)

//HJ107 求解立方根
func HJ107(n float64) {
	e := 0.0001
	t := n
	for math.Abs(t*t*t-n) > e {
		t = t - (t*t*t-n)/(3*t*t)
	}
	fmt.Printf("%0.1f\n", t)
}
