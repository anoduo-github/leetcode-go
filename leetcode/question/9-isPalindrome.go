package question

//9.判断是否是回文数
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x > 0) {
		return false
	}
	//反转一半整数
	res := 0
	for x > res {
		temp := x % 10
		res = res*10 + temp
		x = x / 10
	}
	return res == x || x == res/10
}
