package question

//917. 仅仅反转字母
func reverseOnlyLetters(S string) string {
	//思路，左右指针，一个从左往右，一个从右往左，直到左下标大于等于右小标，期间为字母就相互交换
	//获取长度
	s := []byte(S)
	l := len(s)
	//定义左右指针
	left, right := 0, l-1
	for left < right {
		//判断是否是字母
		if !isEng(s[left]) {
			left++
			continue
		}
		if !isEng(s[right]) {
			right--
			continue
		}
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return string(s)
}

func isEng(b byte) bool {
	if b >= byte('a') && b <= byte('z') {
		return true
	} else if b >= byte('A') && b <= byte('Z') {
		return true
	} else {
		return false
	}
}
