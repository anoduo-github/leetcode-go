package graph

import "sort"

//克鲁斯卡尔算法

//AB14 最小生成树
func miniSpanningTree(n int, m int, cost [][]int) int {
	fa := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		fa[i] = i
	}
	var find func(a int) int
	find = func(x int) int {
		if x != fa[x] {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}
	sort.Slice(cost, func(i, j int) bool {
		return cost[i][2] < cost[j][2]
	})
	ans := 0
	for i := 0; i < m; i++ {
		if find(cost[i][0]) != find(cost[i][1]) {
			ans += cost[i][2]
			fa[find(cost[i][0])] = fa[find(cost[i][1])]
		}
	}
	return ans
}
