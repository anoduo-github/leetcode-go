package question

//79. 单词搜索
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	res := false
	var recv func(i, j, k int)
	recv = func(i, j, k int) {
		if i == -1 || i == len(board) || j == -1 || j == len(board[0]) || visited[i][j] {
			return
		}
		if board[i][j] == word[k] {
			if k == len(word)-1 {
				res = true
				return
			}
			k++
			visited[i][j] = true
		} else {
			return
		}
		//递归找左，右，上，下边的
		recv(i, j-1, k)
		recv(i, j+1, k)
		recv(i-1, j, k)
		recv(i+1, j, k)
		visited[i][j] = false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			recv(i, j, 0)
			if res {
				return true
			}
		}
	}

	return res
}
