package question

//5. 最长回文子串
func longestPalindrome(s string) string {
	max := 1
	res := s[0:1]
	for i := 0; i < len(s)-max; i++ {
		temp := len(s) - 1 //初始指向尾部
		index := i         //头指针
		last := temp       //尾指针
		for index < last {
			if s[index] == s[last] {
				//相等则继续比较
				index++
				last--
			} else {
				//不等则重置头尾指针，重来
				temp--
				index = i
				last = temp
			}
		}
		//跳出循环表示比较完毕
		if max < temp-i+1 {
			max = temp - i + 1
			res = s[i : temp+1]
		}
	}
	return res
}
