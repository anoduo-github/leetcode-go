package question

//202. 快乐数
func isHappy(n int) bool {
	fn := func(num int) int {
		res := 0
		for ; num > 0; num /= 10 {
			i := num % 10
			res += i * i
		}
		return res
	}

	m := make(map[int]bool)
	for ; n != 1 && !m[n]; n, m[n] = fn(n), true {
	}
	return n == 1
}
