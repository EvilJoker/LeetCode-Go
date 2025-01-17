package leetcode

/*
* 思路：

使用 map 记录前一个index,

遍历数组当命中时，检查 index 之差是否在 k 内，不是就返回。

*
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) <= 1 {
		return false
	}
	if k <= 0 {
		return false
	}
	record := make(map[int]bool, len(nums))
	for i, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}
	return false
}
