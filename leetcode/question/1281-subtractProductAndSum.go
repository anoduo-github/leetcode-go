package question

//1281. 整数的各位积和之差
func subtractProductAndSum(n int) int {
	plus := 1
	sum := 0
	for {
		temp := n % 10
		plus = plus * temp
		sum = sum + temp
		n = n / 10
		if n == 0 {
			break
		}
	}
	return plus - sum
}
