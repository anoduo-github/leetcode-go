package graph

import "fmt"

//图

const (
	FULL = "顶点集合已满,无法插入新顶点"
	OVER = "下标超出矩阵范围"
)

//Graph //矩阵图
type Graph struct {
	vertexList []string //存储顶点数组
	edges      [][]int  //存储图对应的邻结矩阵
	numOfEdges int      //图的边的数目
	index      int      //最后一个顶点的后一位下标
	visited    []bool   //顶点是否被标记的数组
}

//NewGraph 初始化一个大小为n的矩阵图
func NewGraph(n int) *Graph {
	//初始化二维数组切片
	rows := make([][]int, n)
	for i := 0; i < n; i++ {
		rows[i] = make([]int, n)
	}

	return &Graph{
		vertexList: make([]string, n),
		edges:      rows,
		numOfEdges: 0,
		index:      0,
		visited:    make([]bool, n),
	}
}

//InsertVertex 插入一个顶点
func (g *Graph) InsertVertex(vertex string) {
	if g.index-1 > len(g.vertexList) {
		panic(FULL)
	}
	g.vertexList[g.index] = vertex
	g.index++
}

//GetVertexSize 获取顶点数量
func (g *Graph) GetVertexSize() int {
	return g.index
}

//InsertEdges 添加边
func (g *Graph) InsertEdges(row, col, weight int) {
	n := len(g.vertexList)
	if row > n-1 || col > n-1 {
		panic(OVER)
	}
	g.edges[row][col] = weight
	g.edges[col][row] = weight
	g.numOfEdges++
}

//GetNumOfEdges 获取边数目
func (g *Graph) GetNumOfEdges() int {
	return g.numOfEdges
}

//getNextNeighbor 获取前一个邻接结点的下一个邻接结点
func (g *Graph) getNextNeighbor(r, c int) int {
	//表示此时c已经是最大范围了，没有下一个邻接结点了
	if c+1 >= g.index {
		return -1
	}

	//循环查找
	for i := c + 1; i < g.index; i++ {
		if g.edges[r][i] > 0 {
			return i
		}
	}
	return -1
}

//PrintGraph 输出矩阵数组
func (g *Graph) PrintGraph() {
	for i := -1; i < g.index; i++ {
		if i == -1 {
			fmt.Print("  ")
			for j := 0; j < g.index; j++ {
				fmt.Printf("%s ", g.vertexList[j])
			}
			fmt.Print("\n")
		} else {
			fmt.Printf("%s ", g.vertexList[i])
			for k := 0; k < g.index; k++ {
				fmt.Printf("%d ", g.edges[i][k])
			}
			fmt.Print("\n")
		}
	}
}

//dfs 深度优先遍历
func (g *Graph) dfs(i int) {
	//输出第一个节点
	fmt.Print(g.vertexList[i] + "->")
	//将该节点设置为已访问
	g.visited[i] = true

	//查找结点i的第一个邻接结点w
	w := g.getNextNeighbor(i, -1)
	for ; w != -1; w = g.getNextNeighbor(i, w) {
		if !g.visited[w] { //如果结点未被访问，递归
			g.dfs(w)
		}
	}
}

//Dfs 深度优先遍历
func (g *Graph) Dfs() {
	fmt.Println("深度优先遍历")
	for i := 0; i < g.index; i++ {
		if !g.visited[i] {
			g.dfs(i)
		}
	}
}

//bfs 广度优先遍历
func (g *Graph) bfs(i int) {
	var u int
	var w int
	queue := make([]int, 0) //切片代替队列

	//输出第一个结点
	fmt.Print(g.vertexList[i] + "->")
	//将该结点设置为已访问
	g.visited[i] = true
	//将结点下标放入切片
	queue = append(queue, i)

	for len(queue) > 0 {
		//取出切片的第一个下标
		u = queue[0]
		queue = queue[1:] //缩小范围
		//得到第一个邻接结点的下标
		w = g.getNextNeighbor(u, -1)
		for ; w != -1; w = g.getNextNeighbor(u, w) {
			if !g.visited[w] { //未被访问
				fmt.Print(g.vertexList[w] + "->")
				g.visited[w] = true
				queue = append(queue, w)
			}
		}
	}
}

//Bfs 广度优先遍历
func (g *Graph) Bfs() {
	fmt.Println("广度优先遍历")
	for i := 0; i < g.index; i++ {
		if !g.visited[i] {
			g.bfs(i)
		}
	}
}

//TestGraph 测试图
func TestGraph() {
	g := NewGraph(8)
	g.InsertVertex("1")
	g.InsertVertex("2")
	g.InsertVertex("3")
	g.InsertVertex("4")
	g.InsertVertex("5")
	g.InsertVertex("6")
	g.InsertVertex("7")
	g.InsertVertex("8")

	g.InsertEdges(0, 1, 1)
	g.InsertEdges(0, 2, 1)
	g.InsertEdges(1, 3, 1)
	g.InsertEdges(1, 4, 1)
	g.InsertEdges(3, 7, 1)
	g.InsertEdges(4, 7, 1)
	g.InsertEdges(2, 5, 1)
	g.InsertEdges(2, 6, 1)
	g.InsertEdges(5, 6, 1)

	g.PrintGraph()

	//g.Dfs()

	g.Bfs()
}
