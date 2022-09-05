package graph

import (
	"fmt"
)

//AB13 【模板】拓扑排序
func TopologicalSort() {
	//n节点数，m边数
	var n, m int
	fmt.Scanln(&n, &m)
	index := make([]int, n+1)   //节点入度
	out := make([]int, n+1)     //节点出度
	graph := make([][]int, n+1) //二维矩阵，记录节点关系
	for i := 0; i < n+1; i++ {
		graph[i] = make([]int, n+1)
	}
	//构建二维矩阵
	for i := 0; i < m; i++ {
		var begin, end int
		fmt.Scanln(&begin, &end)
		index[end]++
		out[begin]++
		graph[begin][end] = 1
	}
	//存结果
	res := make([]int, 0)
	//创建队列,存符合要求的节点
	queue := make([]int, 0)
	for i := 0; i < n+1; i++ {
		if index[i] == 0 && out[i] > 0 {
			queue = append(queue, i)
		}
	}
	//遍历队列
	for len(queue) > 0 {
		//取出队首，放入res
		temp := queue[0]
		queue = queue[1:]
		res = append(res, temp)
		//遍历找到temp的下一个节点，将它的入度减一
		for i := 0; i < n+1; i++ {
			if i != temp && graph[temp][i] == 1 {
				index[i]--
				if index[i] == 0 {
					queue = append(queue, i)
				}
			}
		}
	}
	if len(res) != n {
		fmt.Println(-1)
	} else {
		for _, v := range res {
			fmt.Print(v)
			fmt.Print(" ")
		}
	}
}
