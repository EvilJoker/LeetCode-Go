package leetcode

/*
题目: 找两个墙之前最大的水

思路：两个指针，不断更新索引，找出最大的乘积

优化：贪心，那边墙低移动那边。墙低的一边移动更有可能变大。
*/
func maxArea(height []int) int {
	maxA, left, right := 0, 0, len(height)-1

	for left < right {

		if height[left] < height[right] {

			maxA = max(maxA, (right-left)*height[left])
			left++

		} else {

			maxA = max(maxA, (right-left)*height[right])
			right--

		}
	}
	return maxA
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
