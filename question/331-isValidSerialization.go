package question

import "strings"

//331. 验证二叉树的前序序列化
func isValidSerialization(preorder string) bool {
	//计数器初始化为1，是因为加入根节点时，根节点只有两个出度没有入度。
	//为了方便处理，故意加上一个入度，这样遍历扫描时就减去一个，计数器就为0了
	count := 1
	arr := strings.Split(preorder, ",")
	for _, v := range arr {
		//表示减去一个入度
		count--
		if count < 0 {
			//假如再遍历完成前小于0，就表示不符合要求
			return false
		}
		if v != "#" {
			//表示增加两个出度
			count += 2
		}
	}
	return count == 0
}

/* 用树的出度和入度解决
一个非空节点有一个入度和两个出度
空节点只有一个入度
根节点只有两个出度
树的出度之和和入度之和相等 */
