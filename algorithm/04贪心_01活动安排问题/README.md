# 贪心之活动安排问题

## 题目介绍

活动安排问题（Activity Selection Problem）是一个经典的贪心算法问题。问题描述如下：

给定一组活动，每个活动都有一个开始时间和结束时间。你需要选择尽可能多的活动，使得它们不会重叠，且每个活动只能选一次。

**输入**：一组活动，每个活动有开始时间和结束时间。

**输出**：选择不重叠的最大活动数。

## 解题思路

该问题可以通过贪心算法解决，核心思想是：**选择结束时间最早的活动**，这样可以保证剩下的时间能够安排更多的活动。 ----> 关键

### 步骤

1. 将活动按照结束时间进行排序。
2. 从结束时间最早的活动开始选择，选择后排除与其重叠的活动。
3. 重复上述步骤，直到所有活动都被遍历。

贪心选择的依据是活动结束时间最早的优先，这样可以留下更多时间给后续活动。

## 示例

假设有以下活动，每个活动由开始时间和结束时间组成：

```text
活动1：开始时间 = 1, 结束时间 = 3
活动2：开始时间 = 2, 结束时间 = 5
活动3：开始时间 = 4, 结束时间 = 6
活动4：开始时间 = 6, 结束时间 = 7
活动5：开始时间 = 5, 结束时间 = 8

活动排序后按照结束时间升序排列：

text
复制代码
活动1：开始时间 = 1, 结束时间 = 3
活动2：开始时间 = 2, 结束时间 = 5
活动3：开始时间 = 4, 结束时间 = 6
活动4：开始时间 = 6, 结束时间 = 7
活动5：开始时间 = 5, 结束时间 = 8
选择活动时，首先选择结束时间为3的活动（活动1），然后选择结束时间为6的活动（活动3），最后选择结束时间为7的活动（活动4）。最终选择了活动1、活动3、活动4。

```

```golang

package main

import (
 "fmt"
 "sort"
)

// Activity 结构体表示一个活动
type Activity struct {
 start, end int
}

// 按照结束时间排序活动
func sortActivities(activities []Activity) {
 sort.Slice(activities, func(i, j int) bool {
  return activities[i].end < activities[j].end
 })
}

// 选择活动
func selectActivities(activities []Activity) []Activity {
 // 按结束时间排序
 sortActivities(activities)

 // 选择活动
 selectedActivities := []Activity{}
 lastEndTime := -1

 for _, activity := range activities {
  if activity.start >= lastEndTime {
   selectedActivities = append(selectedActivities, activity)
   lastEndTime = activity.end
  }
 }

 return selectedActivities
}

func main() {
 // 活动数组
 activities := []Activity{
  {1, 3},
  {2, 5},
  {4, 6},
  {6, 7},
  {5, 8},
 }

 // 选择不重叠的活动
 selectedActivities := selectActivities(activities)

 // 输出选择的活动
 fmt.Println("Selected activities:")
 for _, activity := range selectedActivities {
  fmt.Printf("Start: %d, End: %d\n", activity.start, activity.end)
 }
}

```

## 特点

贪心算法：每次选择当前最优解，贪心策略是选择结束时间最早的活动。
时间复杂度：排序操作的时间复杂度是 O(n log n)，遍历活动的时间复杂度是 O(n)，因此总时间复杂度为 O(n log n)。
空间复杂度：O(n)，主要用于存储活动列表和最终选择的活动。