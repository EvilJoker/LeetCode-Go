package leetcode

/*
思路:
两个字符串的最长公共子序列，可以看成是两个字符串的最长公共子串，因此可以考虑动态规划。
建模：dp [i][j] 分别是 text1[0..i] 和 text2[0..j] 的最长公共子序列的长度
初始化：dp[0][j] = 0, dp[i][0] = 0
状态转移方程：
if text1[i-1] == text2[j-1] 代表 字符匹配，
：则 dp[i][j] = dp[i-1][j-1] + 1
：否则 dp[i][j] = max(dp[i-1][j], dp[i][j-1]) 取最大
*/
func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// 初始化二维数组, 默认是 0 初始值，不用在填充
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 填充数组
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			// 转移方程
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}

		}
	}
	return dp[m][n]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
