package leetcode

/*
	笨办法方法一：本法找下一步的位置，能迈出最远的那个位置，就作为下一次的起点

题目：给定一个非负整数数组，最初位于数组的第一个位置。数组中的每个元素代表在该位置可以跳跃的最大长度。判断是否能够到达最后一个位置。
思路：

可以跳跃最大，序号 i + a[i] >= len(a), 最后一步的值，大于数组长度

在 i ~ i+a(i) 序号中找 i+a[i] 最大的值。并将 新节点作为 下一次起点
*/
func canJump1(nums []int) bool {
	n, next_jump := len(nums), 0

	for next_jump < n {
		max_jump := next_jump // 最远能够到的范围
		next_tmp := next_jump // 下一次的起点
		// 遍历起点能够到的范围
		for j := next_jump + 1; j <= nums[next_jump]+next_jump && j < n; j++ {
			// 更新
			if max_jump < nums[j]+j {
				max_jump = nums[j] + j
				next_tmp = next_tmp // 最大的一步作为下一次起点
			}

			// 超出数组长，满足条件
			if max_jump > n-2 {
				return true
			}
		}
		// 原地踏步，说明走不出去了
		if next_tmp == next_jump {
			return false
		}
		// 更新下一个节点
		next_jump = next_tmp
	}

	return false

}

/*
方法二：
方法一的缺点是要重复计算，比如中间某一步是最远，那么下一步开始时，就要重新计算

优化思路: 什么条件下，不能到达。
-- 假设当前是 i ,能到达的最大距离是 i + a[i]， 遍历完 i ~ i+ a[i]，发现不大于 i + a[i]
-- 意味着，不能更新新元素的，永远卡死，那就报错
-- 所以维持 max_jump 距离。遍历所有数组，一旦有一次不能到 max_jump 就说明没戏

## 解题思路

- 给出一个非负数组，要求判断从数组 0 下标开始，能否到达数组最后一个位置。
- 这一题比较简单。如果某一个作为 `起跳点` 的格子可以跳跃的距离是 `n`，那么表示后面 `n` 个格子都可以作为 `起跳点`。可以对每一个能作为 `起跳点` 的格子都尝试跳一次，把 `能跳到最远的距离maxJump` 不断更新。如果可以一直跳到最后，就成功了。如果中间有一个点比 `maxJump` 还要大，说明在这个点和 maxJump 中间连不上了，有些点不能到达最后一个位置。
*/

func canJump(nums []int) bool {
	n := len(nums)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}

	max_jump := 0 // 最大射程

	for i, v := range nums {
		// 不在最大射程，说明前面所有遍历没到当前位置
		if i > max_jump {
			return false
		}
		// 更新
		max_jump = max(max_jump, i+v)
	}
	return true

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
