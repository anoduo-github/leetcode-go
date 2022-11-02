package huawei

//HJ71 字符串通配符
func HJ71() {
	/* var reg string
	fmt.Scanln(&reg)
	var s string
	fmt.Scanln(&s)

	res := false //默认失败
	i, j := 0, 0
	for i < len(reg) && j < len(s) {
		if reg[i] == s[j] {
			i++
			j++
			continue
		}
		if reg[i] != s[j] && (reg[i] != '?' || reg[i] != '*') {
			break
		}
		if reg[i] == '?' {
			i++
			j++
		} else if reg[i] == '*' {
			//遍历找到reg下一个不为*的字符，统计里面的*和？数量
			x := 1         //表示*的数量，这里加上i的*
			y := 0         //表示？的数量
			index := i + 1 //i右边的第一个下标
			for index < len(reg) {
				if reg[index] == '*' {
					x++
					index++
				} else if reg[index] == '?' {
					y++
					index++
				} else {
					break
				}
			}
			//遍历s 找到与*后第一个字母相同的字母，且跨度不能小于x+y
			cur := index //跳过中间长度x+y
			for cur < len(s) {
				if s[cur] == reg[index] {
					//判断长度
					cur++
				}
			}
		}
	}
	fmt.Println(res) */
}
