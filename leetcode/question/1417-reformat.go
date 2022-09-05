package question

//1417. 重新格式化字符串
func reformat(s string) string {
	n := make([]rune, 0)
	t := make([]rune, 0)
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			t = append(t, v)
		}
		if v >= '0' && v <= '9' {
			n = append(n, v)
		}
	}
	ln := len(n)
	lt := len(t)
	i := 0
	res := make([]rune, 0)
	if ln > lt {
		if ln-lt > 1 {
			return ""
		}
		for i < lt {
			res = append(res, n[i])
			res = append(res, t[i])
			i++
		}
		res = append(res, n[i])
	} else if ln == lt {
		for i < lt {
			res = append(res, n[i])
			res = append(res, t[i])
			i++
		}
	} else {
		if lt-ln > 1 {
			return ""
		}
		for i < ln {
			res = append(res, t[i])
			res = append(res, n[i])
			i++
		}
		res = append(res, t[i])
	}
	return string(res)
}
