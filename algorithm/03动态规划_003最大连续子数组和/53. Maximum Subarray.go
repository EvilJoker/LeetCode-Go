package leetcode

import "math"

// 解法一 DP
/* 思路：
建模： dp[i]，以当前节点结束的最大子序列和
初始化： dp[0] = 0
状态转移：dp[i] = max(dp[i-1]+nums[i], nums[i]) // 因为是连续，所以要么是当前节点本身，要么是当前节点加上前一个节点
返回值：dp[] 中最大的数值
*/
func maxSubArray(nums []int) int {
	n := len(nums)

	dp, maxSum := make([]int, n+1), math.MinInt32

	for i := 1; i <= n; i++ {
		dp[i] = max(dp[i-1]+nums[i-1], nums[i-1])
		maxSum = max(maxSum, dp[i])
	}

	return maxSum

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
