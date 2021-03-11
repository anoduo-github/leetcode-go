package practice

/* 任给一个数组，其中只有一个元素是单独出现，其他是成对出现，输出单独的元素。
例如： 输入： {2，2，1，1，4，4，7}
输出：7 */

func Pr2(nums []int) int {
	//map
	m := make(map[int]int)
	//遍历
	for _, v := range nums {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	//遍历
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return 0
}
