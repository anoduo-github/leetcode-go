package question

const m int = 1e9 + 7

//1175. 质数排列
func numPrimeArrangements(n int) int {
	//遍历质数索引,求索引和
	count := 0
	for i := 2; i <= n; i++ {
		if isZhi(i) {
			count++
		}
	}
	//非质数索引数
	no := n - count
	//求阶乘
	return jiecheng(count) * jiecheng(no) % m
}

//判断是否是质数
func isZhi(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

//阶乘
func jiecheng(num int) int {
	res := 1
	for i := 2; i <= num; i++ {
		res = res * i % m
	}
	return res
}
