# 回溯法之 N 皇后问题

## 题目介绍

N 皇后问题是经典的回溯法问题之一，目标是在 N×N 的棋盘上摆放 N 个皇后，使得任意两个皇后不在同一行、同一列或同一斜线上。

**输入**：棋盘大小 N。

**输出**：所有满足条件的皇后摆放方案。

## 解题思路

1. 使用回溯法逐行尝试放置皇后。
2. 在每一行中，尝试将皇后放在某一列，并判断是否满足不在同一列和同一斜线。
3. 如果满足条件，则继续尝试下一行；如果不满足条件，则回溯到上一行，尝试其他列。
4. 递归终止条件为所有皇后都已成功放置。

### 回溯步骤

- 初始化棋盘和辅助数组。
- 从第一行开始递归放置皇后：
  - 检查当前列和对角线是否可以放置。
  - 如果可以放置，则标记当前点并继续递归下一行。
  - 如果放置不成功，则撤销标记，回溯到上一步。

## 示例

```golang
package main

import "fmt"

// solveNQueens 解决 N 皇后问题
func solveNQueens(n int) [][]string {
    results := [][]string{}
    board := make([][]byte, n)
    for i := range board {
        board[i] = make([]byte, n)
        for j := range board[i] {
            board[i][j] = '.'
        }
    }

    var backtrack func(row int)
    backtrack = func(row int) {
        if row == n {
            // 转化棋盘为结果
            result := make([]string, n)
            for i := range board {
                result[i] = string(board[i])
            }
            results = append(results, result)
            return
        }

        for col := 0; col < n; col++ {
            // 检查当前列、主对角线、副对角线是否可用
            if !isValid(board, row, col, n) {
                continue
            }
            // 放置皇后
            board[row][col] = 'Q'
            // 递归下一行
            backtrack(row + 1)
            // 撤销当前行的放置
            board[row][col] = '.'
        }
    }

    backtrack(0)
    return results
}

// isValid 判断当前位置是否可以放置皇后
func isValid(board [][]byte, row, col, n int) bool {
    // 检查当前列是否有皇后
    for i := 0; i < row; i++ {
        if board[i][col] == 'Q' {
            return false
        }
    }

    // 检查主对角线
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if board[i][j] == 'Q' {
            return false
        }
    }

    // 检查副对角线
    for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
        if board[i][j] == 'Q' {
            return false
        }
    }

    return true
}

func main() {
    n := 8
    results := solveNQueens(n)
    fmt.Printf("Total Solutions: %d\n", len(results))
    for _, solution := range results {
        for _, line := range solution {
            fmt.Println(line)
        }
        fmt.Println()
    }
}
```
