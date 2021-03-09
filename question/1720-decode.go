package question

//1720. 解码异或后的数组
func decode(encoded []int, first int) []int {
	l := len(encoded)
	res := make([]int, l+1)
	res[0] = first
	for i := 0; i < l; i++ {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}
