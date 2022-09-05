package algorithm

import "fmt"

//物品	重量		价格
//吉他	1			1500
//音响	4			3000
//电脑	3			2000

//表格
//物品	0		1		2		3		4
//		0		0		0		0		0
//吉他	0		1500	1500	1500	1500
//音响	0		1500	1500	1500	3000
//电脑	0		1500	1500	2000	3500

//感觉核心就是通过表格列出大致情况，然后将表格转化为二维数组，寻找规律，精简出公式

//KnapsackProblemExample 动态规划例子（背包问题）
func KnapsackProblemExample() {
	w := []int{1, 4, 3}            //物品的重量
	val := []int{1500, 3000, 2000} //物品的价值
	c := 4                         //背包的容量
	n := 3                         //物品的个数

	//创建二维数组v[i][j]表示前i个物品占用j容量的最大价值,对应表格，所以范围是（n+1,c+1）
	//初始化
	v := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		v[i] = make([]int, c+1)
	}

	//创建一个二维数组，记录放入商品的情况,并初始化
	p := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = make([]int, c+1)
	}

	//遍历执行
	for i := 1; i < len(v); i++ {
		for j := 1; j < len(v[i]); j++ {
			if w[i-1] > j { //如果当前第i个物品的重量大于背包容量,注意i指的是二维数组v的下标，i-1对应的是数组w
				v[i][j] = v[i-1][j] //表示当前的价值还是等于上一次的价值，因为新的大于背包容量没有放进去嘛
			} else {
				if v[i-1][j] < val[i-1]+v[i-1][j-w[i-1]] {
					v[i][j] = val[i-1] + v[i-1][j-w[i-1]]
					p[i][j] = 1
				} else {
					v[i][j] = v[i-1][j]
				}
			}
		}
	}

	//打印情况
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			fmt.Printf("%d ", v[i][j])
		}
		fmt.Print("\n")
	}

	fmt.Println("***********************")

	//输出如何放的
	row := len(p) - 1
	col := len(p[0]) - 1
	for row > 0 && col > 0 {
		if p[row][col] == 1 {
			fmt.Printf("第%d个物品放入了背包\n", row)
			col -= w[row-1]
		}
		row--
	}
}
