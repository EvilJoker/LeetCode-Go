package leetcode

import (
	"math"
	"sort"
)

// 解法一 O(n^2)
/*
和三数之和和类型，区别是三数之和要求列出所有的组合（去重）。
这道题是返回最接近的值。

定义 diff 值去夹逼
a---b---c， 遍历 a (相同值就滑动)

计算 a+b+c 之和
// 更新
当 b<c 之时
更新 diff 值。

// 移动 bc 指针

当 sum ==0 返回。
否则继续滑动。

这样就得到 最接近的 bc


*/

func threeSumClosest(nums []int, target int) int {
	// 赋值： diff 是最小的差值， res 是最终结果
	n, res, diff := len(nums), 0, math.MaxInt32
	if n > 2 {
		sort.Ints(nums)
		for i := 0; i < n-2; i++ {
			// 当 index 相同时 就滑动
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			j, k := i+1, n-1 // 赋值，都在i 的右边取值
			for j < k {

				sum := nums[i] + nums[j] + nums[k]
				// 不断更新记录最小的值
				if abs(sum-target) < diff {
					res, diff = sum, abs(sum-target)
				}
				if sum == target {
					// 一旦相等继续滑动
					return res
				} else if sum > target {
					k--
				} else {
					j++
				}
			}
		}
	}
	return res
}

// 解法二 暴力解法 O(n^3)
func threeSumClosest1(nums []int, target int) int {
	res, difference := 0, math.MaxInt16
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if abs(nums[i]+nums[j]+nums[k]-target) < difference {
					difference = abs(nums[i] + nums[j] + nums[k] - target)
					res = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return res
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
