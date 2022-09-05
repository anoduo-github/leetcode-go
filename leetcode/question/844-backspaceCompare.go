package question

//844. 比较含退格的字符串
func backspaceCompare(s string, t string) bool {
	return bp(s) == bp(t)
}

func bp(s string) string {
	res := make([]byte, 0)
	arr := []byte(s)
	for _, v := range arr {
		if v == '#' {
			if len(res) > 0 {
				res = res[0 : len(res)-1]
			}
		} else {
			res = append(res, v)
		}
	}
	return string(res)
}
