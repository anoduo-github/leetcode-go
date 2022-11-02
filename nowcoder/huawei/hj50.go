package huawei

import "strconv"

//3+2*{1+2*[-4/(8-6)+7]}
//8-0-0-6
//8-0-(0-6)

//HJ50 四则运算
func HJ50(s string) {
	nums := make([]int, 0)
	chars := make([]byte, 0)

	t := 1
	for i, j := 0, 0; j < len(s); j++ {
		if s[j] >= '0' && s[j] <= '9' {
			if j == len(s)-1 {
				n, _ := strconv.Atoi(s[i : j+1])
				nums = append(nums, t*n)
				t = 1
			}
		} else {
			n, _ := strconv.Atoi(s[i:j])
			nums = append(nums, t*n)
			t = 1
			i = j + 1
			c := s[j]
			if c == '(' || c == '[' || c == '{' {
				chars = append(chars, '(')
			} else if c == ')' || c == ']' || c == '}' {

			}
		}
	}
}
