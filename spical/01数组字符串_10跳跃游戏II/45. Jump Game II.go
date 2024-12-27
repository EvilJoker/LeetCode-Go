package leetcode

/*
题目： 最小的步数
给定一个非负整数数组，你最初位于数组的第一个位置。数组中的每个元素代表你在该位置可以跳跃的最大长度。你的目标是使用最少的跳跃次数到达数组的最后一个位置。

思路：

解法：贪心算法
核心思想：

用最少步数跳到数组的最后一个位置。
每一步尽可能跳到能到达的最远位置。

操作： 尝试 1步是否能到达终点，不行就再加，直到到达边界
*/
func jump(nums []int) int {
	n := len(nums)
	if len(nums) < 2 {
		return 0
	}

	jumps := 0      // 跳跃次数
	currentEnd := 0 // 当前区间能到达的最远位置
	farthest := 0   // 全局能到的最远位置

	for i, v := range nums {
		farthest = max(farthest, i+v)

		//跳一步，能到的最远距离。当到达当前区间的末尾时， 还未到数组末尾 就说明。得再加一步。

		if i == currentEnd { // 区间末尾
			jumps++
			currentEnd = farthest // 体现贪心的地方，到

			// 如果当前区间已经覆盖最后一个位置，可以提前退出
			if currentEnd >= n-1 {
				break
			}
		}

	}

	return jumps

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func jump(nums []int) int {
//     if len(nums) == 1 {
//         return 0
//     }
//     needChoose, canReach, step := 0, 0, 0
//     for i, x := range nums {
//         if i+x > canReach {
//             canReach = i + x
//             if canReach >= len(nums)-1 {
//                 return step + 1
//             }
//         }
//         if i == needChoose {
//             needChoose = canReach
//             step++
//         }
//     }
//     return step
// }
