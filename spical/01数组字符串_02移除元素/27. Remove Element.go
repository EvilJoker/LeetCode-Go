package leetcode

/* 思路：保持原有顺序剔除目标值，返回的是剩余元素的个数
只返回个数--- 长度 n - count
保证数据排序，类似插入排序，维护末尾索引，不断添加新元素

*/

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0 // 数组长度

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			// 因为存在 val 值，j 不会发生误覆盖
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
