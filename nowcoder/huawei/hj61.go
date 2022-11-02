package huawei

//HJ61 放苹果
func HJ61(m, n int) int {
	return kinds(m, n)
}

//分两种情况，一种每个都至少放一个，即f(m-n,n)
//另一种至少一个空盘,即f(m,n-1)
//则f(m,n)=f(m,n-1)+f(m-n,n)
//递归
func kinds(m, n int) int {
	if m < 0 || n < 0 {
		return 0
	}
	if m < 2 || n == 1 {
		return 1
	}
	return kinds(m, n-1) + kinds(m-n, n)
}
