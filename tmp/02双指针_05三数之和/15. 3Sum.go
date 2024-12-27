package leetcode

import (
	"sort"
)

/*
/*
最优解：双指针但是有技巧, 需要记忆边界

思路 ： 排序后的数组(升序) A
a-----b-----c， 三数之和为0， 就是找a + b + c = 0 的,
遍历 a(先滑动), 在右边找bc : 可以搜索到所有情况避免重复
当 a+ b+c > 0. c-- 缩小值
当 a+ b+c < 0  b++ 增大值。
当 a+ b +c =0 ，同时b++ c--（先滑动）

重要其他：结果的去重
思路1： 笨办法，无脑添加，最后比较去重
思路2： a,b,c 发生滑动式，如果下一个值相同，就使劲滑动。注意滑动的时候注意边界
*/

func threeSum(nums []int) [][]int {
	// 结果存储
	result := [][]int{}

	// 首先对数组进行排序，便于去重和双指针查找
	sort.Ints(nums)

	// 遍历数组，每次固定一个数 nums[i]
	for i := 0; i < len(nums)-2; i++ {

		// 跳过重复的数字，避免结果重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 使用双指针查找其余两个数
		left, right := i+1, len(nums)-1

		// 当左指针小于右指针时进行查找
		for left < right {
			// 计算三数之和
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// 如果找到三数之和为 0 的组合，则加入结果集
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// 跳过重复的 left 和 right
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 移动双指针，继续寻找其他组合
				left++
				right--
			} else if sum < 0 {
				// 如果和小于 0，说明左指针的值过小，需要右移增大
				left++
			} else {
				// 如果和大于 0，说明右指针的值过大，需要左移减小
				right--
			}
		}
	}

	// 返回结果集
	return result
}
