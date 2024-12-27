package leetcode

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。注意你不能在买入股票前卖出股票。
Input: [7,1,5,3,6,4]
    Output: 5
思路：
遍历数组维护目前为止最小 的值 min
当前值减去 min， 就是 目前为止最大利润


*/

func maxProfit(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	// 维护最小价格，最大利润，动态更新
	minPrice, maxProfit := prices[0], 0

	for i := 1; i < len(prices); i++ {
		// 目前为止最大利润
		currentProfit := prices[i] - minPrice
		maxProfit = max(maxProfit, currentProfit)
		minPrice = min(minPrice, prices[i])

	}
	return maxProfit
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
