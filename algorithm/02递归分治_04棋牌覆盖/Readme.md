# 定义

递归分治法解决棋牌覆盖问题：  
递归分治是一种将大问题分解为多个相同性质的小问题，并通过递归求解的方法。在棋牌覆盖问题中，我们需要在一个棋盘上放置特定的棋子，按照问题的规则覆盖棋盘的所有空格。

例如：使用“L”形骨牌覆盖一个 2^n × 2^n 的棋盘，棋盘上有一个位置是特殊空白（如障碍或起始点）。

## 示例

```golang
package main

import "fmt"

// 定义棋盘结构
type ChessBoard struct {
 board [][]int
 size  int
}

var tileID int // 骨牌编号

// 创建棋盘
func NewChessBoard(size int) *ChessBoard {
 board := make([][]int, size)
 for i := range board {
  board[i] = make([]int, size)
 }
 return &ChessBoard{board: board, size: size}
}

// 覆盖棋盘
func (cb *ChessBoard) Cover(x, y, specialX, specialY, size int) {
 if size == 1 {
  return // 基本情况：棋盘为 1×1 时，无需操作
 }

 tileID++
 currentTile := tileID
 half := size / 2

 // 左上子棋盘
 if specialX < x+half && specialY < y+half {
  // 特殊点在左上
  cb.Cover(x, y, specialX, specialY, half)
 } else {
  // 左上无特殊点，用 L 形骨牌覆盖
  cb.board[x+half-1][y+half-1] = currentTile
  cb.Cover(x, y, x+half-1, y+half-1, half)
 }

 // 右上子棋盘
 if specialX < x+half && specialY >= y+half {
  // 特殊点在右上
  cb.Cover(x, y+half, specialX, specialY, half)
 } else {
  // 右上无特殊点，用 L 形骨牌覆盖
  cb.board[x+half-1][y+half] = currentTile
  cb.Cover(x, y+half, x+half-1, y+half, half)
 }

 // 左下子棋盘
 if specialX >= x+half && specialY < y+half {
  // 特殊点在左下
  cb.Cover(x+half, y, specialX, specialY, half)
 } else {
  // 左下无特殊点，用 L 形骨牌覆盖
  cb.board[x+half][y+half-1] = currentTile
  cb.Cover(x+half, y, x+half, y+half-1, half)
 }

 // 右下子棋盘
 if specialX >= x+half && specialY >= y+half {
  // 特殊点在右下
  cb.Cover(x+half, y+half, specialX, specialY, half)
 } else {
  // 右下无特殊点，用 L 形骨牌覆盖
  cb.board[x+half][y+half] = currentTile
  cb.Cover(x+half, y+half, x+half, y+half, half)
 }
}

// 打印棋盘
func (cb *ChessBoard) PrintBoard() {
 for _, row := range cb.board {
  fmt.Println(row)
 }
}

func main() {
 n := 3 // 棋盘大小为 2^n × 2^n
 size := 1 << n
 specialX, specialY := 2, 3 // 特殊点坐标

 chessBoard := NewChessBoard(size)
 tileID = 0 // 初始化骨牌编号
 chessBoard.Cover(0, 0, specialX, specialY, size)
 chessBoard.PrintBoard()
}
```

## 特点

分治思想：将问题分解为更小的子问题，递归求解，适合处理具有重复子结构的问题。
递归基：当棋盘大小为 1×1 时，停止递归。
时间复杂度：O(n^2)，因为每次递归分割棋盘为四部分，总体覆盖每个单元格一次。
空间复杂度：O(log(n))，递归栈的深度取决于棋盘的大小。
