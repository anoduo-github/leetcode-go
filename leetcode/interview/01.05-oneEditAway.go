package interview

//面试题 01.05. 一次编辑

func oneEditAway(first string, second string) bool {
	f := []rune(first)
	s := []rune(second)
	d := len(f) - len(s)
	count := 0

	if d == 1 {
		//表示少一位,删除
		i, j := 0, 0
		for i < len(f) && j < len(s) {
			if f[i] == s[j] {
				i++
				j++
			} else {
				i++
				count++
			}
		}
		return count <= 1
	}

	if d == -1 {
		//表示多一位，插入
		i, j := 0, 0
		for i < len(f) && j < len(s) {
			if f[i] == s[j] {
				i++
				j++
			} else {
				j++
				count++
			}
		}
		return count <= 1
	}

	if d == 0 {
		//表示替换
		i := 0
		for i < len(f) {
			if f[i] != s[i] {
				count++
			}
			i++
		}
		return count <= 1
	}

	return false
}
