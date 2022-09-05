package question

import "strings"

//1108. IP 地址无效化
func defangIPaddr(address string) string {
	arr := strings.Split(address, ".")
	ip := arr[0] + "[.]" + arr[1] + "[.]" + arr[2] + "[.]" + arr[3]
	return ip
}
