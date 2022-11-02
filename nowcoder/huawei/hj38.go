package huawei

import "fmt"

//HJ38 求小球落地5次后所经历的路程和第5次反弹的高度
func HJ38() {
	var n int
	fmt.Scanln(&n)

	var res float64
	height := float64(n)
	for i := 0; i < 5; i++ {
		res += height
		height /= float64(2)
		if i < 4 {
			res += height
		}
	}
	fmt.Println(res)
	fmt.Printf("%0.5f\n", height)
}
