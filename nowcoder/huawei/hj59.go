package huawei

import "fmt"

//找出字符串中第一个只出现一次的字符

func hj59() {
	ss := make([]string, 0)
	for {
		var s string
		fmt.Scanln(&s)
		if len(s) == 0 {
			break
		}
		ss = append(ss, s)
	}
	for _, v := range ss {
		m := make(map[rune]int)
		count := 0
		for _, s := range v {
			if _, ok := m[s]; ok {
				m[s]++
			} else {
				m[s] = 1
			}
		}
		for _, s := range v {
			if m[s] == 1 {
				count++
				fmt.Println(string(s))
				break
			}
		}
		if count == 0 {
			fmt.Println(-1)
		}
	}
}
