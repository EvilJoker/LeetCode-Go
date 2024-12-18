# 定义

二分查找是一种在有序数组中查找元素的算法。它通过每次将查找范围减半来快速缩小搜索范围。二分查找的时间复杂度为 O(log n)，比线性查找（O(n)）更高效。

## 示例

golang

```golang
package main

import "fmt"

// 二分查找函数
func binarySearch(arr []int, target int) int {
    left, right := 0, len(arr)-1
    
    // 不断缩小查找范围
    for left <= right {
        mid := left + (right-left)/2

        // 找到目标元素
        if arr[mid] == target {
            return mid
        }

        // 如果目标比中间元素小，缩小右边界
        if arr[mid] > target {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    // 未找到目标元素
    return -1
}

func main() {
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
    target := 7
    result := binarySearch(arr, target)
    
    if result != -1 {
        fmt.Printf("Element found at index: %d\n", result)
    } else {
        fmt.Println("Element not found")
    }
}
```

## 特点

- 时间复杂度：O(log n)，适用于大规模数据。
- 空间复杂度：O(1)，仅需常数空间，使用迭代的方式。
- 适用场景：只能用于已排序的数组或列表。
- 优点：比线性查找（O(n)）效率高，尤其对于大数据集来说。
- 缺点：要求输入数据必须是有序的，否则无法使用二分查找。
