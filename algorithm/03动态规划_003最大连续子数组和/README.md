# 最大子数组和

## 题目介绍

给定一个整数数组 `nums`，找到具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

### 示例

**输入：** `nums = [-2,1,-3,4,-1,2,1,-5,4]`  
**输出：** `6`  
**解释：** 子数组 `[4,-1,2,1]` 的和最大，为 `6`。

**输入：** `nums = [1]`  
**输出：** `1`

---

## 解题思路

### 方法：动态规划

1. 定义一个变量 `currentSum` 表示当前子数组的和。
2. 定义一个变量 `maxSum` 表示全局最大和。
3. 遍历数组，逐步更新 `currentSum`：
   - 如果 `currentSum` 小于 0，则抛弃当前子数组，重新开始新的子数组；
   - 否则，继续累加当前元素。
4. 每次更新 `maxSum` 为 `currentSum` 和 `maxSum` 中的较大值。
5. 最终返回 `maxSum`。

### 算法复杂度

- **时间复杂度：** O(n)，遍历数组一次。
- **空间复杂度：** O(1)，只使用常量额外空间。

---

## demo

```golang
package main

import (
 "fmt"
 "math"
)

// maxSubArray 求最大子数组和
func maxSubArray(nums []int) int {
 // 初始化最大和和当前和
 maxSum := math.MinInt32 // 全局最大和，初始为最小整数
 currentSum := 0         // 当前子数组和

 for _, num := range nums {
  // 如果当前和为负数，舍弃，重新开始新的子数组
  if currentSum < 0 {
   currentSum = 0
  }
  // 累加当前数字到当前和
  currentSum += num
  // 更新最大和
  if currentSum > maxSum {
   maxSum = currentSum
  }
 }
 return maxSum
}

func main() {
 // 示例测试用例
 nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
 fmt.Println("最大子数组和为：", maxSubArray(nums)) // 输出：6
}
```

## 特点

适用场景

在数组或序列中寻找连续子数组的极值问题。
优点

时间复杂度低，O(n)，适合处理大规模数组。
逻辑简单，代码简洁易懂。
扩展应用

可以进一步扩展到二维数组的最大子矩阵和问题。
