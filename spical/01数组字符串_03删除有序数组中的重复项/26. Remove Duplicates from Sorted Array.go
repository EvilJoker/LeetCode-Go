package leetcode

// 思路：有序数组的重复项，如果重复那就是连续重复。
// 类似移除元素，维护新数组索引，如果值和上一个相等，那就是认为是无效值，值不增加
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 1 //新数组索引，至少一个值
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			// 值有效时，直接赋值不重复的索引，丢弃重复的值
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
