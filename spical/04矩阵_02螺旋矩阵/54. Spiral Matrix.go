package leetcode

/*
题目：顺时针螺旋访问矩阵
思路：
1. 标识符控制方向
2. 遇到边界就收缩边界
*/
func spiralOrder(matrix [][]int) []int {
	row, col := len(matrix), len(matrix[0])

	board := []int{col - 1, row - 1, 0, 0} // 右下左上
	direct := 0                            // 0123 右下左上

	ret := []int{}
	i, j := 0, 0 // 起始点
	// 结束点：
	count := 0

	for count < row*col {
		count++
		// fmt.Printf("{%d, %d} e%d  c %d  dir %d, board %v \n", i, j, 0, count, direct, board)
		switch direct {
		case 0:
			ret = append(ret, matrix[i][j])
			// 到达右边界
			if j == board[0] {
				board[3] += 1
				direct = 1
				i++ // 更新起始点
				continue
			}
			j++

		case 1:
			ret = append(ret, matrix[i][j])
			// 到达下边界
			if i == board[1] {
				board[0] -= 1
				direct = 2
				j-- // 更新起始点
				continue
			}
			i++

		case 2:
			ret = append(ret, matrix[i][j])
			// 到达左边界
			if j == board[2] {
				board[1] -= 1
				direct = 3
				i-- // 更新起始点
				continue
			}
			j--
		case 3:
			ret = append(ret, matrix[i][j])
			// 到达上边界
			if i == board[3] {
				board[2] += 1
				direct = 0
				j++ // 更新起始点
				continue
			}
			i--

		}
	}
	return ret

}

// 解法 1
func spiralOrder3(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}
	if len(matrix) == 1 {
		for i := 0; i < len(matrix[0]); i++ {
			res = append(res, matrix[0][i])
		}
		return res
	}
	if len(matrix[0]) == 1 {
		for i := 0; i < len(matrix); i++ {
			res = append(res, matrix[i][0])
		}
		return res
	}
	visit, m, n, round, x, y, spDir := make([][]int, len(matrix)), len(matrix), len(matrix[0]), 0, 0, 0, [][]int{
		{0, 1},  // 朝右
		{1, 0},  // 朝下
		{0, -1}, // 朝左
		{-1, 0}, // 朝上
	}
	for i := 0; i < m; i++ {
		visit[i] = make([]int, n)
	}
	visit[x][y] = 1
	res = append(res, matrix[x][y])
	for i := 0; i < m*n; i++ {
		x += spDir[round%4][0]
		y += spDir[round%4][1]
		if (x == 0 && y == n-1) || (x == m-1 && y == n-1) || (y == 0 && x == m-1) {
			round++
		}
		if x > m-1 || y > n-1 || x < 0 || y < 0 {
			return res
		}
		if visit[x][y] == 0 {
			visit[x][y] = 1
			res = append(res, matrix[x][y])
		}
		switch round % 4 {
		case 0:
			if y+1 <= n-1 && visit[x][y+1] == 1 {
				round++
				continue
			}
		case 1:
			if x+1 <= m-1 && visit[x+1][y] == 1 {
				round++
				continue
			}
		case 2:
			if y-1 >= 0 && visit[x][y-1] == 1 {
				round++
				continue
			}
		case 3:
			if x-1 >= 0 && visit[x-1][y] == 1 {
				round++
				continue
			}
		}
	}
	return res
}

// 解法 2
func spiralOrder2(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return nil
	}

	n := len(matrix[0])
	if n == 0 {
		return nil
	}

	// top、left、right、bottom 分别是剩余区域的上、左、右、下的下标
	top, left, bottom, right := 0, 0, m-1, n-1
	count, sum := 0, m*n
	res := []int{}

	// 外层循环每次遍历一圈
	for count < sum {
		i, j := top, left
		for j <= right && count < sum {
			res = append(res, matrix[i][j])
			count++
			j++
		}
		i, j = top+1, right
		for i <= bottom && count < sum {
			res = append(res, matrix[i][j])
			count++
			i++
		}
		i, j = bottom, right-1
		for j >= left && count < sum {
			res = append(res, matrix[i][j])
			count++
			j--
		}
		i, j = bottom-1, left
		for i > top && count < sum {
			res = append(res, matrix[i][j])
			count++
			i--
		}
		// 进入到下一层
		top, left, bottom, right = top+1, left+1, bottom-1, right-1
	}
	return res
}
