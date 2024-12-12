package leetcode

/*
/*
思路：利用分治法实现快速排序，通过选择一个基准值，将数组分成小于基准值和大于基准值的两部分，递归排序，最终得到有序数组。
平均时间复杂度：O(n log n)
*/

func sort(nums []int) []int {
	result := append([]int(nil), nums...) // 创建副本保证不修改底层数据

	quicksort(result, 0, len(nums)-1)

	return result

}

func quicksort(nums []int, start, end int) {
	// 终止条件
	if start >= end {
		return
	}
	// 分割
	p := partion(nums, start, end)
	// 递归
	quicksort(nums, 0, p)
	quicksort(nums, p+1, end)
}

// 返回分割字符串
func partion(nums []int, start, end int) int {

	// 使用第一个字符串分割数组
	// 类似插入排序，当发现数字小于分隔符时，就扩展数组长度，---- 只是不排序
	// 放置的的技巧：index 代表当前小于 key的数组 ,最后一个位置(预留的，始终为空)。 新数字进来时
	// 1. 新数字放置到 inex 的位置 ---> 新数字位置空出来
	// 2. 扩展 数组长度，---> index + 1, 会侵占下一个数字，放置到 新数字位置
	// 3. index + 1 就变成了新的空位

	index, key := start, nums[start]
	for i := start + 1; i < end+1; i++ {
		if nums[i] < key {
			nums[index], nums[i] = nums[i], nums[index+1]
			index++
		}
	}
	// 回填，此时 index 左边小，右边大
	nums[index] = key
	return index

}
