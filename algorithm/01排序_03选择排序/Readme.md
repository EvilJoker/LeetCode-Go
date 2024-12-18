# 定义

**选择排序**: 选择排序是一种简单直观的排序算法。它的基本思想是每次从未排序的部分中选择最小（或最大）的元素，放到已排序部分的末尾，依次完成整个排序过程。

## 示例

```golang
package main

import "fmt"

// 选择排序实现
func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        // 假设当前索引的元素是最小的
        minIndex := i
        for j := i + 1; j < n; j++ {
            // 找到更小的元素
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        // 交换当前索引和最小值索引的元素
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("原始数组:", arr)
    selectionSort(arr)
    fmt.Println("排序后数组:", arr)
}
```

## 输出

原始数组: [64 25 12 22 11]
排序后数组: [11 12 22 25 64]

## 特点

- **时间复杂度**: O(n²)，无论数据是否有序，每次都要比较剩余未排序的元素。  
- **空间复杂度**: O(1)，原地排序，不需要额外空间。  
- **稳定性**: 不稳定，交换操作可能破坏相同元素的相对顺序。  
- **适用场景**: 数据量小且对性能要求不高时使用。  
