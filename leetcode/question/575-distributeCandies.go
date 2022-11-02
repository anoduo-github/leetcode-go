package question

//575. 分糖果
func distributeCandies(candyType []int) int {
	m := make(map[int]struct{})
	count := 0
	n := len(candyType) / 2
	for _, v := range candyType {
		if _, ok := m[v]; !ok {
			count++
			m[v] = struct{}{}
			if count == n {
				break
			}
		}
	}
	return count
}
