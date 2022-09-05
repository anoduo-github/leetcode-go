package question

import "math"

/* 动态规划
abcde abc

x 0 1 2 3 4 5	i
0 0 0 0 0 0 0
1 0 1 1 1 1 1
2 0 1 1 2 2 2
3 0 1 1 2 2 3

j

x=y
m[i][j]=m[i-1][j-1]+1

x!=y
m[i][j]=max(m[i][j-1],m[i-1][j])
*/

//1143. 最长公共子序列
func longestCommonSubsequence(text1 string, text2 string) int {
	//动态规划
	//初始化一个二维数组，m[i][j]表示text1的前i个字符和text2的前j个字符的最大公共子序列长度
	//+1是因为，确保m[i][0]=m[0][j]=0
	m := make([][]int, len(text1)+1)
	for i := 0; i <= len(text1); i++ {
		m[i] = make([]int, len(text2)+1)
	}
	//遍历两个字符串
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				m[i+1][j+1] = m[i][j] + 1
			} else {
				m[i+1][j+1] = int(math.Max(float64(m[i+1][j]), float64(m[i][j+1])))
			}
		}
	}
	return m[len(text1)][len(text2)]
}
