# 贪心之最优装载问题

## 题目介绍

最优装载问题（Optimal Loading Problem）是一个经典的贪心算法问题。问题描述如下：

假设你有一个卡车，卡车有一个最大承载重量，你需要将一些物品装上卡车，每个物品有一定的重量和价值。你需要选择一些物品装入卡车，使得装载的物品总重量不超过卡车的承载限制，同时最大化装载的物品的总价值。

**输入**：

- 卡车的最大承载重量 `W`。
- 每个物品的重量和价值的列表。

**输出**：

- 选择的物品的最大价值。

## 解题思路

该问题可以通过贪心算法解决，核心思想是：**每次选择单位价值最高的物品**，直到装不下更多的物品。

### 步骤

1. 计算每个物品的单位价值（即：物品的价值除以重量）。
2. 按照单位价值从高到低对物品进行排序。
3. 选择单位价值最高的物品，将它们装入卡车，直到卡车无法再装下更多物品。

贪心选择的依据是单位价值最高的物品，因为它能够在同样的重量下提供更高的价值。

## 示例

假设有以下物品，每个物品有重量和价值：

| 物品 | 重量 | 价值 |
|------|------|------|
| 物品1 | 2    | 3    |
| 物品2 | 3    | 4    |
| 物品3 | 4    | 5    |
| 物品4 | 5    | 8    |

卡车最大承载重量为 5。

### 解法

1. 计算每个物品的单位价值：
   - 物品1：3 / 2 = 1.5
   - 物品2：4 / 3 ≈ 1.33
   - 物品3：5 / 4 = 1.25
   - 物品4：8 / 5 = 1.6
2. 按照单位价值排序：物品4 > 物品1 > 物品2 > 物品3
3. 选择物品4和物品1，装载重量为 5，最大价值为 3 + 8 = 11。

## Golang 示例

```golang
package main

import (
 "fmt"
 "sort"
)

// Item 结构体表示一个物品
type Item struct {
 weight, value int
}

// ByValuePerWeight 用于按照物品的单位价值排序
type ByValuePerWeight []Item

func (items ByValuePerWeight) Len() int {
 return len(items)
}

func (items ByValuePerWeight) Less(i, j int) bool {
 // 按单位价值从高到低排序
 return float64(items[i].value)/float64(items[i].weight) > float64(items[j].value)/float64(items[j].weight)
}

func (items ByValuePerWeight) Swap(i, j int) {
 items[i], items[j] = items[j], items[i]
}

// 选择物品装载卡车，返回最大价值
func knapsack(W int, items []Item) int {
 // 按照单位价值从高到低排序物品
 sort.Sort(ByValuePerWeight(items))

 totalValue := 0
 totalWeight := 0

 for _, item := range items {
  // 如果当前物品能装下，加入卡车
  if totalWeight+item.weight <= W {
   totalValue += item.value
   totalWeight += item.weight
  }
 }

 return totalValue
}

func main() {
 // 物品数组
 items := []Item{
  {weight: 2, value: 3},
  {weight: 3, value: 4},
  {weight: 4, value: 5},
  {weight: 5, value: 8},
 }

 // 卡车最大承载重量
 W := 5

 // 计算最大价值
 maxValue := knapsack(W, items)

 // 输出最大价值
 fmt.Println("Max value:", maxValue)
}
```
