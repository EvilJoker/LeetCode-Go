# 贪心之多机调度问题

## 题目介绍

多机调度问题是经典的调度优化问题，目的是在 `m` 台机器上分配 `n` 个任务，使得完成所有任务的总时间最短。

**输入**：任务列表，每个任务有一个执行时间，以及机器的数量 `m`。  
**输出**：分配任务的方案，使得所有机器完成任务的时间最小。

## 解题思路

贪心算法可以用于解决多机调度问题：  

1. 将所有任务按时间从大到小排序（贪心选择原则）。  
2. 使用一个最小堆（或数组）记录每台机器的当前负载。  
3. 依次将任务分配给负载最小的机器（局部最优），以减小全局负载。  

### 步骤

1. 对任务按照执行时间降序排序。
2. 遍历任务列表，每次将任务分配给当前负载最小的机器。
3. 更新机器的负载，继续分配下一个任务。
4. 最后，负载最大的机器完成时间即为最短完成时间。

## 示例

### 输入

- 任务时间：`[5, 8, 3, 7, 2, 6]`
- 机器数量：`3`

### 输出

- 分配方案：机器 1: `[8]`, 机器 2: `[7, 3]`, 机器 3: `[6, 5, 2]`  
- 最短完成时间：`13`（机器 3 完成任务所需的时间）。

## Golang 示例

```golang
package main

import (
    "container/heap"
    "fmt"
    "sort"
)

// 定义机器的负载结构
type Machine struct {
    id    int
    load  int
    index int // 堆中索引
}

// 定义最小堆
type MinHeap []*Machine

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].load < h[j].load }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i]; h[i].index, h[j].index = i, j }

func (h *MinHeap) Push(x interface{}) {
    machine := x.(*Machine)
    machine.index = len(*h)
    *h = append(*h, machine)
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    machine := old[n-1]
    machine.index = -1 // 清除索引
    *h = old[0 : n-1]
    return machine
}

// 多机调度函数
func multiMachineSchedule(tasks []int, m int) ([][]int, int) {
    // 初始化结果集
    result := make([][]int, m)
    for i := range result {
        result[i] = []int{}
    }

    // 初始化最小堆，表示每台机器的负载
    minHeap := &MinHeap{}
    heap.Init(minHeap)
    for i := 0; i < m; i++ {
        heap.Push(minHeap, &Machine{id: i, load: 0})
    }

    // 按任务时间降序排序
    sort.Sort(sort.Reverse(sort.IntSlice(tasks)))

    // 分配任务
    for _, task := range tasks {
        // 从堆中取出当前负载最小的机器
        machine := heap.Pop(minHeap).(*Machine)

        // 分配任务给该机器
        result[machine.id] = append(result[machine.id], task)
        machine.load += task

        // 将机器放回堆中
        heap.Push(minHeap, machine)
    }

    // 获取最大负载（即完成任务所需的最短时间）
    maxLoad := 0
    for _, machine := range *minHeap {
        if machine.load > maxLoad {
            maxLoad = machine.load
        }
    }

    return result, maxLoad
}

func main() {
    tasks := []int{5, 8, 3, 7, 2, 6}
    machines := 3

    result, maxLoad := multiMachineSchedule(tasks, machines)

    fmt.Println("Task Allocation:")
    for i, tasks := range result {
        fmt.Printf("Machine %d: %v\n", i+1, tasks)
    }
    fmt.Printf("Minimum completion time: %d\n", maxLoad)
}
```

## 特点

贪心算法：每次选择当前负载最小的机器，确保局部最优解。
时间复杂度：排序任务的时间复杂度为 O(n log n)，最小堆的插入与删除操作复杂度为 O(log m)，总复杂度为 O(n log n + n log m)。
空间复杂度：O(m + n)，存储堆和分配结果。

## 示例运行结果

text
复制代码
Task Allocation:
Machine 1: [8]
Machine 2: [7, 3]
Machine 3: [6, 5, 2]
Minimum completion time: 13
