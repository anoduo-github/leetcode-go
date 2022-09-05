package offer

//剑指 Offer 10- I. 斐波那契数列
func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	one := 0
	two := 1
	res := 0
	for i := 2; i <= n; i++ {
		res = (one + two) % 1000000007
		one = two
		two = res
	}
	return res
}
