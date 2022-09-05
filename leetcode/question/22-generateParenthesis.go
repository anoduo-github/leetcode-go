package question

//22. 括号生成
func generateParenthesis(n int) []string {
	res := make([]string, 0)
	var recv func(l, r int, s string)
	recv = func(l, r int, s string) {
		if len(s) == 2*n {
			//如果长度等于2n，表示该s满足条件
			res = append(res, s)
			return
		}
		//先添加(
		if l > 0 {
			recv(l-1, r, s+"(")
		}
		//再添加)
		if l < r {
			//用l<r是要保证左括号比右括号数少，也就是用掉的左括号要比右括号多，才能添加右括号
			recv(l, r-1, s+")")
		}
	}
	//开始递归
	recv(n, n, "")
	return res
}
