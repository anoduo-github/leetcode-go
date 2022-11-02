package huawei

import (
	"fmt"
)

//HJ36 字符串加密
func HJ36() {
	var key string
	fmt.Scanln(&key)
	var s string
	fmt.Scanln(&s)

	//构建基本的字母表
	base := make([]byte, 26)
	for i := 0; i < 26; i++ {
		base[i] = byte('a' + i)
	}

	//给key按顺序排序，同时去重
	k := make([]byte, 0) //存新字母表
	m := make(map[byte]struct{})
	for i := 0; i < len(key); i++ {
		if _, ok := m[key[i]]; !ok {
			k = append(k, key[i]) //先存密钥字母
			m[key[i]] = struct{}{}
		}
	}
	for _, v := range base {
		if _, ok := m[v]; !ok {
			k = append(k, v) //再存正常字母表
		}
	}
	arr := []byte(s)
	for i := 0; i < len(arr); i++ {
		arr[i] = k[arr[i]-'a']
	}
	fmt.Println(string(arr))
}
