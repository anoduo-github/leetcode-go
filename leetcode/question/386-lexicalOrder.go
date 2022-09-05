package question

//386. 字典序排数
func lexicalOrder(n int) []int {
	ans := make([]int, 0)
	temp := 1
	for i := 0; i < n; i++ {
		ans = append(ans, temp)
		if temp*10 <= n {
			temp *= 10
		} else {
			for temp%10 == 9 || temp >= n {
				temp /= 10
			}
			temp++
		}
	}
	return ans
}
