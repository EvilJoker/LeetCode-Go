package leetcode

/*
题目：
环形公路上有 n 个加油站， 有gas[] 和 cost[] 两个数组，代表当前能加到的汽油。和到下一站消耗的汽油。
请问是否能跑完全程，能的话，起点是什么。

思路：
总汽油数 >= 总消耗数，就能跑全程.

问题是求起点
1. 暴力求解，挨个试节点。
2. 优化，当某个节点没汽油了，表示什么。---> 之前的节点作为起点都不可行， 起点肯定在后面的节点，所以设下一个点为起点。
3. 起点肯定是 有剩余的

*/

func canCompleteCircuit(gas []int, cost []int) int {

	totalGas, totalCost, start, n := 0, 0, 0, len(gas) // 起点

	for i := 0; i < n; i++ {
		totalGas += gas[i]
		totalCost += cost[i]

		// 油气不够，设置新起点.不能用总的判断。假设某个节点有很多gas ，其他没有，但是这里还是会判断通过。不更新起点
		if cost[i] > gas[i] {
			start = i + 1
		}

	}
	if totalCost > totalGas {
		return -1
	}

	return start

}
