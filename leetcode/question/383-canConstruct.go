package question

//383. 赎金信
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for _, v := range []byte(magazine) {
		m[v]++
	}
	for _, v := range []byte(ransomNote) {
		m[v]--
		if m[v] < 0 {
			return false
		}
	}
	return true
}
