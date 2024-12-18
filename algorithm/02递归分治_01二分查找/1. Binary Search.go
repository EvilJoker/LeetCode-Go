package leetcode

/*
思路：二分查找是一种在有序数组中查找元素的算法。它通过每次将查找范围减半来快速缩小搜索范围。二分查找的时间复杂度为 O(log n)，比线性查找（O(n)）更高效。
*/

// // 二分查找函数, 返回索引
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	// 不断缩小查找范围
	for left <= right {
		mid := left + (right-left)/2

		// 找到目标元素
		if arr[mid] == target {
			return mid
		}

		// 如果目标比中间元素小，缩小右边界
		if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 未找到目标元素
	return -1
}
