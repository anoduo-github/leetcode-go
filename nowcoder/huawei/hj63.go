package huawei

import (
	"fmt"
)

//HJ63 DNA序列
func HJ63() {
	var s string
	fmt.Scanln(&s)
	var n int
	fmt.Scanln(&n)

	//先判断前n个字符长的
	max := 0
	var res string
	count := 0
	for i := 0; i < n; i++ {
		if s[i] == 'C' || s[i] == 'G' {
			count++
		}
	}
	if count > max {
		max = count
		res = s[0:n]
	}

	//每一轮窗口里的最大值
	m := max
	//滑动
	l, r := 1, n
	for r < len(s) {
		temp := m
		//左边界外的第一个
		if s[l-1] == 'C' || s[l-1] == 'G' {
			temp -= 1
		}
		//右边界，右移的那一个
		if s[r] == 'C' || s[r] == 'G' {
			temp += 1
		}
		//判断
		m = temp
		if temp > max {
			max = temp
			res = s[l : r+1]
		}
		l++
		r++
	}
	fmt.Println(res)
}
