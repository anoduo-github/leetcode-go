package algorithm

import "fmt"

//ViolenceMatch 暴力匹配
func ViolenceMatch(str1, str2 string) int {
	//转换成切片
	temp1 := []rune(str1)
	temp2 := []rune(str2)
	length1 := len(temp1) //被匹配字串长度
	length2 := len(temp2) //子串长度
	if length2 > length1 {
		return -1
	}

	i := 0       //i表示temp1的左下标
	j := length2 //j表示temp1的右下标，temp1[i:j]正好是一个temp2的长度

	for j <= length1 {
		if string(temp1[i:j]) == str2 {
			return i
		}
		i++
		j++
	}

	return -1
}

//kmpNext 获取一个字符串的部分匹配值表
func kmpNext(s string) []int {
	arr := []rune(s)
	n := len(arr)

	//初始化n个0
	//res[0]=0,表示长度为1的部分匹配值为0
	res := make([]int, n)

	for i, j := 1, 0; i < n; i++ {
		for j > 0 && arr[i] != arr[j] {
			j = res[j-1]
		}
		if arr[i] == arr[j] {
			j++
		}
		res[i] = j
	}

	return res
}

//kmpSearch kmp匹配
func kmpSearch(s1, s2 string, next []int) int {
	arr1 := []rune(s1)
	arr2 := []rune(s2)
	for i, j := 0, 0; i < len(arr1); i++ {
		for j > 0 && arr1[i] != arr2[j] {
			j = next[j-1]
		}
		if arr1[i] == arr2[j] {
			j++
		}
		if j == len(arr2) {
			return i - j + 1
		}
	}
	return -1
}

//TestKmp 测试kmp
func TestKmp() {
	str1 := "BBC ABCDAB ABCDABCDABDE"
	str2 := "ABCDABD"

	res := kmpNext(str2)
	fmt.Println(res)

	index := kmpSearch(str1, str2, res)
	fmt.Println(index)
}
