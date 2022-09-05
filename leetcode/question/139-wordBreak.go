package question

//139. 单词拆分
func wordBreak(s string, wordDict []string) bool {
	//表示字符串前i个字符组成的单词是否在wordDict中存在
	dp := make([]bool, len(s)+1)
	dp[0] = true

	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[s[j:i]] {
				//表示前j个存在（不含j），从j到i的字符也存在
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
