package question

import "strings"

//796. 旋转字符串
func rotateString(s string, goal string) bool {
	//两个字符串长度必须相等，且s+s表示了s旋转的所有可能，只要s+s包含goal就证明可以
	return len(s) == len(goal) && strings.Contains(s+s, goal)
}
