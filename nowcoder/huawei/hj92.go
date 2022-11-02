package huawei

import "fmt"

//HJ92 在字符串中找出连续最长的数字串
func HJ92(s string) {
	max := 1
	res := make([]string, 0)
	for i := 0; i < len(s); i++ {
		j := i
		for j < len(s) {
			if s[j] >= '0' && s[j] <= '9' {
				j++
			} else {
				break
			}
		}
		if max < j-i {
			max = j - i
			res = make([]string, 0)
			res = append(res, s[i:j])
		} else if max == j-i {
			res = append(res, s[i:j])
		}
		i = j
	}
	for _, v := range res {
		fmt.Print(v)
	}
	fmt.Printf(",%d\n", max)
}
