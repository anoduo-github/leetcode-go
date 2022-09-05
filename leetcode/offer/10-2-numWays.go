package offer

//剑指 Offer 10- II. 青蛙跳台阶问题
func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	i, j := 1, 2
	res := 0
	for k := 3; k <= n; k++ {
		res = (i + j) % 1000000007
		i = j
		j = res
	}
	return res
}
