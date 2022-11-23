package huawei

import "fmt"

//dp[i][j]表示前i个物品预算为j时最大的价值，以此题例：dp[3][1000]=500,可以理解为在预算为1000时，我买了三个物品，价值500
//动态转移方程：dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i]]+v[i])
//解释，w[i]表示第i个物品的费用，v[i]表示第i个物品的价值
//表示两者情况，比如i=3,j=1000，表示预算为1000的情况下，最大值要么是i=2,j=1000的情况，即预算不变，我只买2个的最大价值
//或者是我买了第i=3个，假设第3个物品的费用为400，那么dp[3][1000]=dp[2][600]+v[3],即用预算在600时2个物品的最大值加上第三个物品的价值
//不可能是dp[2][1000]+v[3],因为第3个也有费用，加了第三个，预算就不止1000了

//加上附件
//四种情况
//1 主件	2 主件+附件1	3 主件+附件2	4 主件+附件1+附件2
//将w[i]改为w[i][k],表示第i个物品的第k种情况的费用
//将v[i]改为v[i][k]，表示第i个物品的第k种情况的价值
//k为0到3，即上述的四种情况
//即：dp[i][j]=max(dp[i-1][j],dp[i-1][j-w[i][k]]+v[i][k])

func HJ16() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	n = n / 10
	w := make([][]int, m)
	for i := 0; i < m; i++ {
		w[i] = make([]int, 4)
	}
	t := make([][]int, m)
	for i := 0; i < m; i++ {
		t[i] = make([]int, 4)
	}
	count := 0 //表示主件数
	for i := 0; i < m; i++ {
		var v, p, q int
		fmt.Scan(&v)
		fmt.Scan(&p)
		fmt.Scan(&q)
		if q == 0 {
			w[count][0] = v
			t[count][0] = p
			count++
		} else {
			if w[q-1][1] == 0 {
				w[q-1][1] = w[q-1][0] + v
				t[q-1][1] = t[q-1][0] + p
			} else {
				w[q-1][2] = w[q-1][0] + v
				t[q-1][2] = t[q-1][0] + p

				w[q-1][3] = w[q-1][1] + w[q-1][2] - w[q-1][0]
				t[q-1][3] = t[q-1][1] + t[q-1][2] - t[q-1][0]
			}
		}
	}

	dp := make([][]int, count+1)
	for i := 0; i <= count; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < count; i++ {
		for j := 1; j < n+1; j++ {
			m := 0
			for k := 0; k < 4; k++ {
				var temp int
				if j-w[i][k] > 0 {
					temp = max(dp[i][j], dp[i][j-w[i][k]]+t[i][k])
				} else {
					temp = dp[i][j]
				}
				if temp > m {
					m = temp
				}
			}
			dp[i+1][j] = m
		}
	}
	fmt.Println(dp[count][n])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
