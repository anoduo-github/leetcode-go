package offer

import "strings"

//剑指 Offer 58 - I. 翻转单词顺序
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	i := 0
	arr := make([]string, 0)
	for i < len(s) {
		if s[i] == ' ' {
			i++
			continue
		}
		j := i
		for j < len(s) {
			if s[j] != ' ' {
				j++
				if j == len(s) {
					arr = append(arr, s[i:j])
					i = j + 1
					break
				}
			} else {
				arr = append(arr, s[i:j])
				i = j + 1
				break
			}
		}
	}
	res := ""
	for i := len(arr) - 1; i >= 0; i-- {
		res += arr[i]
		if i > 0 {
			res += " "
		}
	}
	return res
}
