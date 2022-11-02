package huawei

import "fmt"

//HJ52 计算字符串的编辑距离
func HJ52() {
	var src string
	fmt.Scanln(&src)
	var dst string
	fmt.Scanln(&dst)
	var s [26]int
	for _, v := range src {
		s[v-'a']++
	}
	var d [26]int
	for _, v := range dst {
		d[v-'a']++
	}
	a := 0 //dst相对src多的字母数
	b := 0 //dst相对src少的字母数
	for i := 0; i < 26; i++ {
		if s[i] > d[i] {
			b += s[i] - d[i]
		} else if s[i] < d[i] {
			a += d[i] - s[i]
		}
	}
	if a > b {
		//多的补少的，剩下的删除
		fmt.Println(a)
	} else {
		fmt.Println(b)
	}
}
