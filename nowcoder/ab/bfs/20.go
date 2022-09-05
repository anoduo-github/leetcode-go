package bfs

//AB20 走迷宫

var graph [][]int //存图
var sx, sy int    //存起始点x,y
var ex, ey int    //存终点x,y
var queue [][]int //当队列

func MinStep() int {
	n, m := 14, 14
	graph = make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, m)
	}
	sx, sy = 11, 1
	ex, ey = 3, 3
	s := []string{
		"...**.*.......",
		"*.....**.....*",
		"...*..........",
		"*..*.*..*.*...",
		"......**.....*",
		"..*..*.*......",
		"........*...*.",
		"...*.**....*..",
		".*.......*..**",
		".*..*..*.....*",
		"......**......",
		"**..........*.",
		".**.....*.....",
		"...........*..",
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s[i][j] == '*' {
				graph[i][j] = 1
			} else {
				graph[i][j] = 0
			}
		}
	}
	if graph[sx-1][sy-1] == 1 || graph[ex-1][ey-1] == 1 {
		return -1
	}
	queue = make([][]int, 0)
	node := []int{0, sx - 1, sy - 1}
	queue = append(queue, node)
	graph[sx-1][sy-1] = -1
	steps := bfs(n, m)
	return steps
}

func bfs(n, m int) int {
	for len(queue) > 0 {
		//弹出队首
		temp := queue[0]
		if len(queue) == 1 {
			queue = make([][]int, 0)
		} else {
			queue = queue[1:]
		}
		step := temp[0]
		x := temp[1]
		y := temp[2]
		if x == ex && y == ey {
			return step
		}
		if x+1 < n && graph[x+1][y] == 0 {
			//表示可以向下走
			node := []int{step + 1, x + 1, y}
			queue = append(queue, node)
			graph[x+1][y] = -1
		}
		if y+1 < m && graph[x][y+1] == 0 {
			//表示可以向右走
			node := []int{step + 1, x, y + 1}
			queue = append(queue, node)
			graph[x][y+1] = -1
		}
		if x-1 >= 0 && graph[x-1][y] == 0 {
			//表示可以向上走
			node := []int{step + 1, x - 1, y}
			queue = append(queue, node)
			graph[x-1][y] = -1
		}
		if y-1 >= 0 && graph[x][y-1] == 0 {
			//表示可以向左走
			node := []int{step + 1, x, y - 1}
			queue = append(queue, node)
			graph[x][y-1] = -1
		}
	}
	return -1
}
