package huawei

import (
	"fmt"
	"math"
)

//对正浮点数四舍五入
func hj6() {
	var f float64
	fmt.Scanln(&f)
	i, j := math.Modf(f)
	if j >= 0.5 {
		i += 1
	}
	fmt.Println(i)
}
