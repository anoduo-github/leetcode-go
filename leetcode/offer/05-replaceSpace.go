package offer

import "strings"

//剑指 Offer 05. 替换空格
func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
