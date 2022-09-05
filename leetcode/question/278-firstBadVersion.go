package question

import sss "sort"

func isBadVersion(version int) bool {
	if version < 6 {
		return true
	}
	return false
}

//278. 第一个错误的版本
func firstBadVersion(n int) int {
	/* index := 1
	last := n
	for index <= last {
		mid := index + (last-index)/2
		if isBadVersion(mid) {
			last = mid
		} else {
			index = mid + 1
		}
	}
	return index */

	return sss.Search(n, func(version int) bool { return isBadVersion(version) })
}
