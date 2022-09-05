package question

import "strings"

//648. 单词替换
func replaceWords(dictionary []string, sentence string) string {
	//切割句子
	arr := strings.Split(sentence, " ")
	//遍历
	res := ""
	for _, k := range arr {
		temp := k
		for _, v := range dictionary {
			if len(temp) > len(v) && len(k) >= len(v) && k[:len(v)] == v {
				temp = k[:len(v)]
			}
		}
		res += temp + " "
	}
	return res[:len(res)-1]
}
