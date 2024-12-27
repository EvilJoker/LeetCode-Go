package leetcode

/*
思路：利用map 做到 O(n),
*/

func twoSum(nums []int, target int) []int {

	cacheMap := map[int]int{} // key 是值，value 是索引
	for i, v := range nums {
		left := target - v

		// 尝试获取对应的值，存在时 就返回
		if j, ok := cacheMap[left]; ok {
			return []int{i, j}
		}
		// 不存在就放入
		cacheMap[v] = i

	}

	return nil

}

func twoSumbk(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}
