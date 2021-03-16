package huawei

import "fmt"

//正整数A和正整数B 的最小公倍数是指 能被A和B整除的最小的正整数值，设计一个算法，求输入A和B的最小公倍数。

func hj108() {
	var num1 int
	var num2 int
	fmt.Scanf("%d %d\n", &num1, &num2)
	res := num1 * num2
	for i := num1*num2 - 1; i > 0; i-- {
		if i%num1 == 0 && i%num2 == 0 {
			res = i
		}
	}
	fmt.Println(res)
}
