package question

//547. 省份数量
func findCircleNum(isConnected [][]int) int {
	res := 0
	for i := 0; i < len(isConnected); i++ {
		//队列存点
		queue := make([]int, 0)
		//逐行遍历
		for j := 0; j < len(isConnected[0]); j++ {
			if isConnected[i][j] == 1 {
				//当存在连接点，则将该点放入队列
				queue = append(queue, j)
				//并将其值置0，表示已经访问过
				isConnected[i][j] = 0
				isConnected[j][i] = 0
			}
		}
		if len(queue) > 0 {
			//大于0表示有相连的点
			res++
		}
		//判断这些相连的点又没其他相连的点，广度遍历
		for len(queue) > 0 {
			//取第一个
			temp := queue[0]
			//缩小queue
			queue = queue[1:]
			if temp == i {
				continue
			}
			for k := 0; k < len(isConnected[0]); k++ {
				if isConnected[temp][k] == 1 {
					queue = append(queue, k)
					isConnected[temp][k] = 0
					isConnected[k][temp] = 0
				}
			}
		}
	}
	return res
}
