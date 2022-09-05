package question

import sss "sort"

//374. 猜数字大小
func guessNumber(n int) int {
	return sss.Search(n, func(x int) bool { return guess(x) <= 0 })
}

func guess(num int) int {
	if num == 6 {
		return 0
	} else if num < 6 {
		return 1
	} else {
		return -1
	}
}
