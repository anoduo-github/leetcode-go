package question

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var res []string

//17. 电话号码的字母组合
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res = make([]string, 0)
	group(0, digits, "")
	return res

}

func group(index int, digits, temp string) {
	if index == len(digits) {
		//表示下标已经到最后一个的下一个了，把最后一个的字符串加入
		res = append(res, temp)
	} else {
		//取值
		v := phoneMap[string(digits[index])]
		//遍历
		for i := 0; i < len(v); i++ {
			//递归
			group(index+1, digits, temp+string(v[i]))
		}
	}
}
