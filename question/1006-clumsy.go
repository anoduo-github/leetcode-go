package question

//1006. 笨阶乘
func clumsy(N int) int {
	if N == 1 {
		return N
	}
	//数栈
	num := make([]int, 0)
	num = append(num, N)
	N--
	//符号栈,1表示乘，2表示除，3表示加，4表示减
	chs := make([]int, 0)
	//定义符号位
	tag := 1
	for N > 0 {
		switch tag {
		case 1:
			chs = append(chs, 1)
			num = append(num, N)
			tag++
		case 2:
			chs = append(chs, 2)
			num = append(num, N)
			tag++
		case 3:

			tag++
		case 4:
			tag = 1
		}
		N--
	}
	return 0
}
