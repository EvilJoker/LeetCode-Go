package leetcode

// 思路：有序数组的重复项，如果重复那就是连续重复。
// 维护新数组索引，如果值和上一个相等，那就是认为是无效值，值不增加
// 进阶：判断条件更新，绝对不能重复3次

func removeDuplicates(nums []int) int {
	// 2以内直接返回
	if len(nums) < 3 {
		return len(nums)
	}

	j := 2 // 至少是2

	for i := 2; i < len(nums); i++ {

		if nums[i] != nums[i-2] { // 意味着只要不连续三个重复
			nums[j] = nums[i]
			j++
		}

	}
	return j
}
