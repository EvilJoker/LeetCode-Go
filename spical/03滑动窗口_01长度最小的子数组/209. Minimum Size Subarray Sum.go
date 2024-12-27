package leetcode

/* 题目：连续子数组之和 > targent， 且找到最短的 子数组. 否则返回 0

思路：滑动窗口，头尾指针遍历。
当 < target，右指针 右移。
当 > targent，左指针 右移
*/

func minSubArrayLen(target int, nums []int) int {
	left, end, sum, n, minx := 0, 0, 0, len(nums), 0

	for left <= end && end < n {
		if sum >= target {

			minx = min(minx, end-left+1)
			sum -= nums[left]
			left++
			continue
		} else {
			// 右移
			sum += nums[end]
			end++
		}

	}
	return minx

}

func minSubArrayLen1(target int, nums []int) int {
	left, sum, res := 0, 0, len(nums)+1
	for right, v := range nums {
		sum += v
		for sum >= target {
			res = min(res, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
