package question

//709. 转换成小写字母
func toLowerCase(str string) string {
	arr := []byte(str)
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 'A' && arr[i] <= 'Z' {
			arr[i] = arr[i] + 'a' - 'A'
		}
	}
	return string(arr)
}
