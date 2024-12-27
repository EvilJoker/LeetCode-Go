package leetcode

/*
思路： 找出众数,大于 n/2 的元素，数组总是存在众数
方式1： 使用 hash 表 o/n
*/
func majorityElement(nums []int) int {
	maps := make(map[int]int)
	for _, num := range nums {
		maps[num]++
		if maps[num] > len(nums)/2 {
			return num
		}
	}
	return 0
}
