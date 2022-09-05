package question

//1523. 在区间范围内统计奇数数目
func countOdds(low int, high int) int {
	l := low % 2
	h := high % 2
	if low == high {
		if l == 1 {
			return 1
		} else {
			return 0
		}
	}
	if l == 1 && h == 1 {
		//两个都是奇数
		return (high-low-1)/2 + 2
	} else {
		//一个奇数一个偶数,或者两个偶数
		return (high-low-1)/2 + 1
	}
}

/* 7 8 9 10 11 12 13
6-1=5
5/2=2
2+2=4

7 8 9 10 11 12
12-7-1=4
4/2=2
2+1=3

8 9 10 11 12 13
13-8-1=4
4/2=2
2+1=3

8 9 10 11 12
12-8-1=3
3/2=1 */
