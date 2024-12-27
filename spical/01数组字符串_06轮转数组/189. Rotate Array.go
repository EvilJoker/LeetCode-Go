package leetcode

/*
思路：
轮转数组，有技巧，
整体翻转 12345 --> 54321
前k 翻转 54321 -->  45321
后n-k 翻转  45321 --> 45123
*/
func rotate(nums []int, k int) {
	k %= len(nums) // 求余
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])

}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
