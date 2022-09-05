package question

//557. 反转字符串中的单词 III
func reverseWords(s string) string {
	arr := []rune(s)
	i, j := 0, 0
	for j < len(arr) {
		if arr[j] == ' ' {
			index := i
			last := j - 1
			for index < last {
				arr[index], arr[last] = arr[last], arr[index]
				index++
				last--
			}
			i = j + 1
		}
		j++
	}
	j--
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}
