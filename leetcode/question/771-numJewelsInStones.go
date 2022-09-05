package question

//771. 宝石与石头
func numJewelsInStones(jewels string, stones string) int {
	if len(jewels) == 0 || len(stones) == 0 {
		return 0
	}
	//采用map
	//1.先将石头放进map,key为值，value为重复值的数量
	res := make(map[rune]int)
	for _, v := range stones {
		if _, ok := res[v]; ok {
			res[v]++
		} else {
			res[v] = 1
		}
	}
	//2.遍历jewels查询宝石
	sum := 0
	for _, v := range jewels {
		if i, ok := res[v]; ok {
			sum = sum + i
		}
	}
	return sum
}
