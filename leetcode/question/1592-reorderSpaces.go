package question

//1592. 重新排列单词间的空格
func reorderSpaces(text string) string {
	arrs := make([]string, 0)
	count := 0
	i := 0
label:
	for i < len(text) {
		if text[i] == ' ' {
			count++
			i++
		} else {
			j := i + 1
			for ; j < len(text); j++ {
				if text[j] == ' ' {
					count++
					arrs = append(arrs, text[i:j])
					i = j + 1
					continue label
				}
			}
			arrs = append(arrs, text[i:j])
			i = j
		}
	}
	res := ""
	if len(arrs) == 1 {
		res += arrs[0]
		for i := 0; i < count; i++ {
			res += " "
		}
		return res
	}
	nums1 := count / (len(arrs) - 1)
	nums2 := count % (len(arrs) - 1)
	space1 := ""
	space2 := ""
	for i := 0; i < nums1; i++ {
		space1 += " "
	}
	for i := 0; i < nums2; i++ {
		space2 += " "
	}
	for i := 0; i < len(arrs); i++ {
		if i == len(arrs)-1 {
			res += arrs[i] + space2
		} else {
			res += arrs[i] + space1
		}
	}
	return res
}
