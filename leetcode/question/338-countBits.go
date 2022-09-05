package question

//338. 比特位计数
func countBits(n int) []int {
	countOne := func(i int) int {
		count := 0
		for ; i > 0; i &= i - 1 {
			count++
		}
		return count
	}

	res := make([]int, n+1)
	for i := range res {
		res[i] = countOne(i)
	}
	return res
}
