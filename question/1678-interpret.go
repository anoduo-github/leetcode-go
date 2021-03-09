package question

//1678. 设计 Goal 解析器
func interpret(command string) string {
	res := ""
	i := 0
	for i < len(command) {
		if string(command[i]) == "G" {
			res = res + "G"
			i++
		} else if string(command[i]) == "(" {
			if string(command[i+1]) == ")" {
				res = res + "o"
				i = i + 2
			} else {
				res = res + "al"
				i = i + 4
			}
		}
	}
	return res
}
