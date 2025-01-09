package leetcode

import "strconv"

/*
	横竖 9宫格(0,3,9) 中重复1-9 不能重复

思路： 暴力判断即可
*/
func isValidSudoku(board [][]byte) bool {
	// 判断横
	for i := 0; i < 9; i++ {
		cached := [10]int{} // 全0数组，索引代表 1~9 数字代表次数
		for j := 0; j < 9; j++ {
			c := board[i][j]
			if c == '.' {
				continue
			}
			// 有重复
			if cached[c-'0'] != 0 {
				return false
			}
		}

	}

	// 判断竖
	for i := 0; i < 9; i++ {
		cached := [10]int{} // 全0数组，索引代表 1~9 数字代表次数
		for j := 0; j < 9; j++ {
			c := board[j][i]
			if c == '.' {
				continue
			}
			// 有重复
			if cached[c-'0'] != 0 {
				return false
			}
		}

	}

	// 判断 9宫格

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			// ij 就是9宫格偏移量
			cached := [10]int{}
			// 9宫格数字遍历
			for m := 0; m < 9; m++ {
				for n := 0; n < 9; n++ {
					c := board[m+i*3][n+j*3]

					if c == '.' {
						continue
					}
					// 有重复
					if cached[c-'0'] != 0 {
						return false
					}
				}
			}

		}

	}
	return true
}

// 解法一 暴力遍历，时间复杂度 O(n^3)
func isValidSudoku3(board [][]byte) bool {
	// 判断行 row
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[i][j : j+1]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}
	// 判断列 column
	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[j][i]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}
	// 判断 9宫格 3X3 cell
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]int{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					cellVal := board[ii][jj]
					if string(cellVal) != "." {
						index, _ := strconv.Atoi(string(cellVal))
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}
	return true
}

// 解法二 添加缓存，时间复杂度 O(n^2)
func isValidSudoku1(board [][]byte) bool {
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}
	// 遍历一次，添加缓存
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				num := board[r][c] - '0' - byte(1)
				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}
				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true // r,c 转换到box方格中
			}
		}
	}
	return true
}
