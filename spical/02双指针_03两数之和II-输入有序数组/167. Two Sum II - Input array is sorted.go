package leetcode

/*
题目：两数之和：有序数组，找出是否存在两个数字加起来等于 target

思路： 使用map ，简化到O(n)
*/
func twoSum167(numbers []int, target int) []int {

	// key_index
	maps := make(map[int]int, len(numbers))

	for i, v := range numbers {
		// 另一个数
		left := target - v
		// 找到, 注意 0值和索引会重复，所以 为 i+1
		if maps[left] != 0 {
			return []int{maps[left], i + 1}
		}
		// 没找到加入
		maps[v] = i + 1

	}
	return []int{-1, -1}

}
