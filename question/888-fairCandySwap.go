package question

//888. 公平的糖果棒交换
func fairCandySwap(A []int, B []int) []int {
	//A的和
	sum1 := 0
	//B的和
	sum2 := 0
	//建一个map存放b的值
	b := make(map[int]int)
	//返回的结果
	res := make([]int, 0)
	for _, v1 := range A {
		sum1 += v1
	}
	for _, v2 := range B {
		b[v2] = 1
		sum2 += v2
	}
	//假设A减去x, B减去y,则x和y满足以下关系,y=x-(sum(A)-sum(B))/2,让temp=(sum(A)-sum(B))/2
	temp := (sum1 - sum2) / 2
	//遍历A查找x是否满足上述公式
	for _, v3 := range A {
		if _, ok := b[v3-temp]; ok {
			res = append(res, v3)
			res = append(res, v3-temp)
			//找到一组即可
			break
		}
	}
	return res
}
