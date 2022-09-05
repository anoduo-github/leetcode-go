package offer

//剑指 Offer 48. 最长不含重复字符的子字符串
func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}
	max := 0
	for i := 0; i < n-max; i++ {
		m := make(map[byte]struct{}, 0)
		for j := i; j < n; j++ {
			if _, ok := m[s[j]]; ok {
				temp := j - i
				if max < temp {
					max = temp
				}
				break
			} else {
				if j == n-1 {
					temp := j - i + 1
					if max < temp {
						max = temp
					}
				}
				m[s[j]] = struct{}{}
			}
		}
	}
	return max
}
