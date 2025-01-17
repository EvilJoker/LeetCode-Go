package leetcode

import (
	"strconv"
)

/*
题目： 无重复有序数组 中连续数字 比如 2、3、4 应该被 2~4 代替

思路（无序数组）：
设置列表 每个值是 [a,b]，
1. 如果数字在某个值中，就过
2. 和边界相连，扩展边界
3. 不在就创建 [c,c]

思路（有序数组，更简单）：
如果发现不连续就新加一个


*/

func summaryRanges(nums []int) (ans []string) {

	if len(nums) == 0 {
		return []string{}
	}
	startIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			continue
		}
		// 不相等时
		endIndex := i - 1

		if startIndex == endIndex {
			ans = append(ans, strconv.Itoa(nums[startIndex]))
		} else {
			ans = append(ans, strconv.Itoa(nums[startIndex])+"->"+strconv.Itoa(nums[endIndex]))
		}
		startIndex = i

	}

	if startIndex == len(nums)-1 {
		ans = append(ans, strconv.Itoa(nums[startIndex]))
	} else {
		ans = append(ans, strconv.Itoa(nums[startIndex])+"->"+strconv.Itoa(nums[len(nums)-1]))
	}
	return ans

}

func summaryRanges1(nums []int) (ans []string) {
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}
		s := strconv.Itoa(nums[left])
		if left != i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return
}
