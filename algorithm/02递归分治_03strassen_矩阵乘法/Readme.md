# 定义

Strassen 矩阵乘法是一种基于递归分治思想的矩阵乘法算法，旨在减少矩阵乘法所需的标量乘法运算次数。相比传统方法，Strassen 算法通过拆分矩阵和计算部分矩阵来优化大矩阵的乘法运算。

## 示例

```golang
package main

import (
 "fmt"
)

// splitMatrix 将矩阵分为四个子矩阵
func splitMatrix(matrix [][]int) ([][]int, [][]int, [][]int, [][]int) {
 n := len(matrix) / 2
 a11 := make([][]int, n)
 a12 := make([][]int, n)
 a21 := make([][]int, n)
 a22 := make([][]int, n)

 for i := 0; i < n; i++ {
  a11[i] = matrix[i][:n]
  a12[i] = matrix[i][n:]
  a21[i] = matrix[i+n][:n]
  a22[i] = matrix[i+n][n:]
 }
 return a11, a12, a21, a22
}

// addMatrices 两个矩阵相加
func addMatrices(a, b [][]int) [][]int {
 n := len(a)
 result := make([][]int, n)
 for i := 0; i < n; i++ {
  result[i] = make([]int, n)
  for j := 0; j < n; j++ {
   result[i][j] = a[i][j] + b[i][j]
  }
 }
 return result
}

// subtractMatrices 两个矩阵相减
func subtractMatrices(a, b [][]int) [][]int {
 n := len(a)
 result := make([][]int, n)
 for i := 0; i < n; i++ {
  result[i] = make([]int, n)
  for j := 0; j < n; j++ {
   result[i][j] = a[i][j] - b[i][j]
  }
 }
 return result
}

// strassen 矩阵乘法核心
func strassen(a, b [][]int) [][]int {
 n := len(a)
 if n == 1 {
  return [][]int{{a[0][0] * b[0][0]}}
 }

 a11, a12, a21, a22 := splitMatrix(a)
 b11, b12, b21, b22 := splitMatrix(b)

 m1 := strassen(addMatrices(a11, a22), addMatrices(b11, b22))
 m2 := strassen(addMatrices(a21, a22), b11)
 m3 := strassen(a11, subtractMatrices(b12, b22))
 m4 := strassen(a22, subtractMatrices(b21, b11))
 m5 := strassen(addMatrices(a11, a12), b22)
 m6 := strassen(subtractMatrices(a21, a11), addMatrices(b11, b12))
 m7 := strassen(subtractMatrices(a12, a22), addMatrices(b21, b22))

 c11 := addMatrices(subtractMatrices(addMatrices(m1, m4), m5), m7)
 c12 := addMatrices(m3, m5)
 c21 := addMatrices(m2, m4)
 c22 := addMatrices(subtractMatrices(addMatrices(m1, m3), m2), m6)

 // 合并结果矩阵
 result := make([][]int, n)
 for i := 0; i < n/2; i++ {
  result[i] = append(c11[i], c12[i]...)
  result[i+n/2] = append(c21[i], c22[i]...)
 }
 return result
}

func main() {
 a := [][]int{
  {1, 2},
  {3, 4},
 }
 b := [][]int{
  {5, 6},
  {7, 8},
 }

 result := strassen(a, b)
 fmt.Println("Result:")
 for _, row := range result {
  fmt.Println(row)
 }
}
```

## 特点

分治思想：将矩阵拆分成更小的子矩阵递归计算，适合大规模矩阵。
减少乘法次数：使用 7 次标量乘法代替传统算法的 8 次乘法。
适用性：适合矩阵维度是 2 的幂次的情况，否则需要填充矩阵。
空间复杂度高：需要额外的存储空间保存中间结果。
递归特性：适用于递归环境，但在维度较小时递归会有一定性能开销。
