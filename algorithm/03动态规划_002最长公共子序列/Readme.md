# 定义

最长公共子序列（Longest Common Subsequence，LCS）是指在两个序列中找出最长的子序列（可以不连续），并且这个子序列在两个序列中都出现。不同于最长公共子字符串，LCS 不要求子序列是连续的。

## 题目介绍

给定两个字符串 `text1` 和 `text2`，返回它们的最长公共子序列的长度。如果不存在公共子序列，返回 0。

示例：  
输入：`text1 = "abcde"`, `text2 = "ace"`  
输出：`3`  
解释：最长公共子序列是 `"ace"`，它的长度是 3。

## 解题思路

1. **动态规划**：
   - 创建一个二维数组 `dp`，其中 `dp[i][j]` 表示 `text1` 的前 `i` 个字符和 `text2` 的前 `j` 个字符的最长公共子序列的长度。
   - **递推公式**：
     - 如果 `text1[i-1] == text2[j-1]`，说明字符匹配，则 `dp[i][j] = dp[i-1][j-1] + 1`。
     - 否则，取 `dp[i-1][j]` 和 `dp[i][j-1]` 的较大值，即 `dp[i][j] = max(dp[i-1][j], dp[i][j-1])`。
   - 初始化：`dp[0][...]` 和 `dp[...][0]` 初始化为 0，因为空字符串与任何字符串的公共子序列长度为 0。
2. 遍历完成后，`dp[m][n]`（`m` 和 `n` 分别是 `text1` 和 `text2` 的长度）即为结果。

3. **空间优化（可选）**：
   - 由于每次只需要 `dp[i-1][...]` 和当前行 `dp[i][...]` 的值，可以用一维数组优化空间复杂度。

## 示例

```golang
package main

import (
 "fmt"
)

func longestCommonSubsequence(text1 string, text2 string) int {
 // 获取两个字符串的长度
 m, n := len(text1), len(text2)
 
 // 创建二维 dp 数组
 dp := make([][]int, m+1)
 for i := range dp {
  dp[i] = make([]int, n+1)
 }

 // 填充 dp 数组
 for i := 1; i <= m; i++ {
  for j := 1; j <= n; j++ {
   // 如果字符匹配，继承左上角值加 1
   if text1[i-1] == text2[j-1] {
    dp[i][j] = dp[i-1][j-1] + 1
   } else {
    // 否则取上方或左方的较大值
    dp[i][j] = max(dp[i-1][j], dp[i][j-1])
   }
  }
 }

 // 返回最终结果
 return dp[m][n]
}

// 工具函数：取最大值
func max(a, b int) int {
 if a > b {
  return a
 }
 return b
}

func main() {
 text1 := "abcde"
 text2 := "ace"
 result := longestCommonSubsequence(text1, text2)
 fmt.Printf("The length of the Longest Common Subsequence is: %d\n", result)
}
```

## 特点

时间复杂度：O(m *n)，其中 m 和 n 分别是两个字符串的长度。
空间复杂度：O(m* n)，可以通过使用滚动数组优化为 O(min(m, n))。
动态规划表格用于保存中间结果，避免重复计算。
