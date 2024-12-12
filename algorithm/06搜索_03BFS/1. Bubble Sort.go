package leetcode

/*
思路：冒泡排序通过多次比较和交换相邻元素，将最大值或最小值逐步移动到数组的一端。
平均时间复杂度：O(n^2)，因为每次需要遍历未排序部分的所有元素，比较和交换操作成平方级增长。
*/

func sort(nums []int) []int {
	for i := 0; i < len(nums); i++ { // 交换测试
		for j := 0; j < len(nums)-i-1; j++ { // 每一次都产生最大值，所以去除最后一位
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j] //独特的交换方式
			}
		}
	}
	return nums

}
