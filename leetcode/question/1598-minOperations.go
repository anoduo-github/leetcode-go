package question

//1598. 文件夹操作日志搜集器
func minOperations(logs []string) int {
	res := make([]string, 0)
	for _, v := range logs {
		if v == "./" {
			continue
		}
		if v == "../" {
			if len(res) > 0 {
				res = res[0 : len(res)-1]
			}
		} else {
			res = append(res, v)
		}
	}
	return len(res)
}
