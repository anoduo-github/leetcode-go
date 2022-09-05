package question

import "strings"

//824. 山羊拉丁文
func toGoatLatin(sentence string) string {
	//切割字符串
	arr := strings.Split(sentence, " ")
	length := len(arr)
	//存元音字母
	y := "aAeEiIoOuU"
	//遍历
	res := ""
	for i := 0; i < length; i++ {
		temp := []rune(arr[i])
		first := temp[0]
		//判断是否是元音开头
		if strings.Contains(y, string(first)) {
			res += arr[i] + "ma"
		} else {
			temp = append(temp, temp[0])
			res += string(temp[1:]) + "ma"
		}
		//根据索引加a
		for j := 1; j <= i+1; j++ {
			res += "a"
		}
		res += " "
	}
	res = strings.TrimSpace(res)

	return res
}
