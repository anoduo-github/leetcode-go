package question

//461. 汉明距离
func hammingDistance(x int, y int) int {
	//首先异或，1^0=1,0^0=0,1^1=1
	z := x ^ y
	count := 0
	for z > 0 {
		//每次与一次，则二进制数减一
		z = z & (z - 1)
		count++
	}
	return count
}
