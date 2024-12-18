package leetcode

/*

**选择排序**: 选择排序是一种简单直观的排序算法。它的基本思想是每次从未排序的部分中选择最小（或最大）的元素，放到已排序部分的末尾，依次完成整个排序过程
 */

// 选择排序实现
func sort(nums []int) []int {
	result := append([]int(nil), nums...) // 创建副本保证不修改底层数据

	n := len(result)
	for i := 0; i < n-1; i++ {
		// 假设当前索引的元素是最小的
		minIndex := i
		for j := i + 1; j < n; j++ {
			// 找到更小的元素
			if result[j] < result[minIndex] {
				minIndex = j
			}
		}
		// 交换当前索引和最小值索引的元素
		result[i], result[minIndex] = result[minIndex], result[i]
	}
	return result
}
