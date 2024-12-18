package leetcode

/*
思路：冒泡排序通过多次比较和交换相邻元素，将最大值或最小值逐步移动到数组的一端。
平均时间复杂度：O(n^2)，因为每次需要遍历未排序部分的所有元素，比较和交换操作成平方级增长。

归并排序（Merge Sort）是一种分治法的排序算法。它将数组递归地划分成两个子数组，对这两个子数组分别排序后，再将它们合并为一个有序数组。

*/

import "fmt"

// merge 两个有序子数组
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	// 按照顺序合并两个子数组
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 将剩余的元素添加到结果数组
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// mergeSort 实现归并排序
func mergeSort(arr []int) []int {
	// 递归结束条件
	if len(arr) <= 1 {
		return arr
	}

	// 分割数组
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	// 合并两个有序子数组
	return merge(left, right)
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Unsorted array:", arr)
	sortedArr := mergeSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}
