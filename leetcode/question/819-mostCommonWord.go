package question

//819.最常见的单词
func mostCommonWord(paragraph string, banned []string) string {
	j := 0
	arr := []rune(paragraph)
	m := make(map[string]int) //存单词
	for _, v := range banned {
		m[v] = 0 //表示禁止的词
	}

	temp := ""
	for j < len(arr) {
		if arr[j] >= 'a' && arr[j] <= 'z' {
			temp += string(arr[j])
		} else if arr[j] >= 'A' && arr[j] <= 'Z' {
			arr[j] += 'a' - 'A' //大写转小写
			temp += string(arr[j])
		} else {
			if temp != "" {
				if value, ok := m[temp]; ok {
					if value != 0 {
						m[temp]++
					}
				} else {
					m[temp] = 1
				}
				temp = ""
			}
		}
		j++
	}

	if temp != "" {
		if value, ok := m[temp]; ok {
			if value != 0 {
				m[temp]++
			}
		} else {
			m[temp] = 1
		}
		temp = ""
	}

	res := ""
	max := 0
	for k, v := range m {
		if v > max {
			res = k
			max = v
		}
	}
	return res
}
