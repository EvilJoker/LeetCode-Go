# 定义

对称矩阵: 是一种特殊的方阵，其中矩阵的元素满足 `A[i][j] == A[j][i]` 的特性。可以利用对称性减少存储空间，例如只存储矩阵的下三角或上三角部分。

单链表中的对称矩阵: 使用单链表来存储对称矩阵的元素。一般只存储矩阵的一半数据（上三角或下三角），每个节点包含元素值及其对应的矩阵位置。

## 示例

```golang
package main

import "fmt"

// 定义矩阵节点
type MatrixNode struct {
    row   int
    col   int
    value int
    next  *MatrixNode
}

// 打印对称矩阵链表
func printMatrix(head *MatrixNode) {
    for head != nil {
        fmt.Printf("A[%d][%d] = %d\n", head.row, head.col, head.value)
        head = head.next
    }
}

func main() {
    // 创建链表表示的对称矩阵
    head := &MatrixNode{row: 0, col: 0, value: 1}
    head.next = &MatrixNode{row: 0, col: 1, value: 2}
    head.next.next = &MatrixNode{row: 1, col: 1, value: 3}

    printMatrix(head)
}
```

## 特点

使用链表存储对称矩阵时，通常只存储上三角或下三角的非零元素。
节省存储空间，特别是在矩阵很大且稀疏的情况下。
查找元素时需要遍历链表，时间复杂度较高。