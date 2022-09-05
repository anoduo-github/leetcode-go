package question

//1773. 统计匹配检索规则的物品数量
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	l := len(items)
	j := -1
	sum := 0
	switch ruleKey {
	case "type":
		j = 0
	case "color":
		j = 1
	case "name":
		j = 2
	}
	if j == -1 {
		return 0
	}
	for i := 0; i < l; i++ {
		if items[i][j] == ruleValue {
			sum++
		}
	}
	return sum
}
