package huawei

import "fmt"

//HJ71 字符串通配符
func HJ71() {
	/* var reg string
	fmt.Scanln(&reg)
	var s string
	fmt.Scanln(&s) */

	reg := "h*h*ah**ha*h**h***hha"
	s := "hh"

	//递归

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if i == len(reg) && j == len(s) {
			return true
		} else if i == len(reg) || j == len(s) {
			return false
		}
		//两个字符对应的上，相等或者一个为？
		if reg[i] == '?' {
			if !(s[j] >= '0' && s[j] <= '9') && !(s[j] >= 'a' && s[j] <= 'z') {
				return false
			}
			return dfs(i+1, j+1)
		} else if reg[i] == s[j] {
			return dfs(i+1, j+1)
		} else if reg[i] == '*' {
			//为*，一个都不匹配
			if dfs(i+1, j) {
				return true
			}
			//至少匹配一个，此时要求s[j]不能是其他字符
			if !(s[j] >= '0' && s[j] <= '9') && !(s[j] >= 'a' && s[j] <= 'z') {
				return false
			}
			//为*，匹配一个位
			if dfs(i+1, j+1) {
				return true
			}
			//为*，匹配多个位
			return dfs(i, j+1)
		}
		return false
	}

	fmt.Println(dfs(0, 0))
}
