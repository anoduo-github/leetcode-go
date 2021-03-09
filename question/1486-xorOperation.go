package question

//1486. 数组异或操作
func xorOperation(n int, start int) int {
	res := start
	for i := 1; i < n; i++ {
		res = res ^ (2 * i)
	}
	return res
}
