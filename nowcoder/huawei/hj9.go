package huawei

import "fmt"

//输入一个int型整数，按照从右向左的阅读顺序，返回一个不含重复数字的新的整数。
//保证输入的整数最后一位不是0。
func hj9() {
	var num int
	fmt.Scanln(&num)
	res := 0
	m := make(map[int]struct{})
	for num > 0 {
		temp := num % 10
		if _, ok := m[temp]; ok {
			num = num / 10
			continue
		} else {
			m[temp] = struct{}{}
			res = res*10 + temp
			num = num / 10
		}
	}
	fmt.Println(res)
}
