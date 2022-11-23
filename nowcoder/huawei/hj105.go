package huawei

import "fmt"

//HJ105 记负均正II
func HJ105() {
	count1 := 0
	count2 := 0
	sum := 0
	for {
		var n int
		_, err := fmt.Scanln(&n)
		if err != nil {
			break
		}
		if n < 0 {
			count1++
		} else {
			count2++
			sum += n
		}
	}
	fmt.Println(count1)
	res := 0.0
	if count2 > 0 {
		res = float64(sum) / float64(count2)
	}
	fmt.Printf("%0.1f\n", res)
}
