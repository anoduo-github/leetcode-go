package question

//944. 删列造序

func minDeletionSize(strs []string) int {
	count := 0
	for i := 0; i < len(strs[0]); i++ {
		max := int32(0)
		for _, v := range strs {
			temp := []rune(v)[i]
			if temp >= max {
				max = temp
			} else {
				count++
				break
			}
		}
	}
	return count
}
