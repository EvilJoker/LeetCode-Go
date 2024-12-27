# 回溯法之装载问题

## 题目介绍

装载问题是经典的回溯法问题之一，通常描述为：有一个集装箱船的最大载重为 `C`，有 `n` 个货物，其重量分别为 `w1, w2, ..., wn`。需要装载货物，使得装载的总重量尽可能接近但不超过 `C`。

---- 动态规划更简单

**输入**：

- 最大载重 `C`。
- 货物重量数组 `weights`。

**输出**：

- 最优的装载方案及其总重量。

## 解题思路

回溯法通过递归的方式尝试每一种可能的装载方案。每次递归时，选择当前货物是否装载，最终记录最优解。

### 步骤

1. 初始化当前总重量为 0，最优解为 0。
2. 从第一个货物开始递归：
   - 若当前货物重量加上总重量不超过最大载重，则装载当前货物，递归处理下一个货物。
   - 选择跳过当前货物，递归处理下一个货物。
3. 在递归过程中更新最优解。
4. 输出最优装载方案及其重量。

## 示例

### Golang 实现

```golang
package main

import (
 "fmt"
)

// 回溯法解决装载问题
func backtrack(weights []int, capacity int, currentWeight int, index int, bestWeight *int) {
 // 如果当前总重量超过最优重量，更新最优重量
 if currentWeight > *bestWeight {
  *bestWeight = currentWeight
 }
 // 说明：这里有技巧。每次递归都是找 包含当前节点的最优解。一直到装不下为止


 // 遍历剩余货物
 for i := index; i < len(weights); i++ {
  // 如果装载当前货物不会超过容量，递归装载下一个货物
  if currentWeight+weights[i] <= capacity {
    
   backtrack(weights, capacity, currentWeight+weights[i], i+1, bestWeight)
  }
 }
}

func solveLoadingProblem(weights []int, capacity int) int {
 bestWeight := 0
 backtrack(weights, capacity, 0, 0, &bestWeight)
 return bestWeight
}

func main() {
 // 输入货物重量和船的最大载重
 weights := []int{10, 20, 30, 40, 50}
 capacity := 100

 // 求解装载问题
 result := solveLoadingProblem(weights, capacity)

 // 输出结果
 fmt.Printf("最优装载重量: %d\n", result)
}
```

关键注释说明
backtrack 函数：

使用递归模拟每种装载方案。
参数 currentWeight 表示当前已装载的总重量，index 表示当前正在处理的货物索引。
使用 bestWeight 记录最优装载方案的重量。
剪枝：

如果当前货物装载后会超出最大载重，则跳过该货物，避免不必要的计算。
主函数调用：

初始化 bestWeight 为 0。
递归调用 backtrack 计算最优重量。

## 特点

贪心与回溯结合：通过尝试每种可能方案，找到最优解。
时间复杂度：最坏情况下为 O(2^n)，其中 n 是货物数量，因为每个货物有两种选择：装载或不装载。
