package leetcode

/*
思路
插入排序的核心思想是：通过构建一个有序序列，将未排序的数据逐步插入到已排序部分的适当位置，直到所有数据都完成排序。
它适合小规模数据集，在数据量较小时性能良好。

将前i个元素，当成排序好的数据。后面元素加入时，往前滑动

时间复杂度: O(n²)（因为每插入一个元素需要遍历已排序部分的平均长度为 n/2）
*/

func sort(nums []int) []int {
	result := append([]int(nil), nums...) // 创建副本保证不修改底层数据

	// 遍历数组，从第二个元素开始，因为第一个元素默认已排序
	for i := 1; i < len(result); i++ {

		// num[j] 就是新插入的元素，如果前面的元素小于 自己，就和前面的交换。这样就保证始终是自己在比较。不是说明已经排序到位，退出即可
		for j := i; result[j] < result[j-1] && j > 0; j-- {
			result[j], result[j-1] = result[j-1], result[j]
		}

	}
	return result

}
