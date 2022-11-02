package question

import "sort"

//670. 最大交换
func maximumSwap(num int) int {
	//转换成数组
	temp1 := make([]int, 0)
	for num > 0 {
		temp1 = append(temp1, num%10) //倒着存的
		num /= 10
	}
	length := len(temp1)
	//复制
	temp2 := make([]int, length)
	copy(temp2, temp1)
	//排序,因为是递增排序，所以后面都是倒着比较
	sort.Ints(temp2)
	//比较，找到第一个不一致的
	for i := length - 1; i >= 0; i-- {
		if temp1[i] != temp2[i] {
			//表示找到，但是要注意可能有相同的数，要替换最后一个
			// 1993  9931 此时i=0，但是要替换第二个9
			index := 0
			for j := i - 1; j >= 0; j-- {
				if temp1[j] == temp2[i] {
					index = j
				}
			}
			//交换
			temp1[i], temp1[index] = temp1[index], temp1[i]
			break
		}
	}
	//返回
	res := 0
	for i := length - 1; i >= 0; i-- {
		res = res*10 + temp1[i]
	}
	return res
}
