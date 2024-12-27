package leetcode

/*
题目：
多次买卖。保证利润最大.

思路：
相当于一个波浪线，制取上升部分.

当 p[i] > p[i-1], max += p[i] - p[i-1] + max


*/

func maxProfit122(prices []int) int {
	sum := 0
	for i := 1; i < len(prices); i++ {
		// 只加增长部分利润
		if prices[i]-prices[i-1] > 0 {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}
