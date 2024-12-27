package leetcode

/*
思路：
字符串只有两个前进方向，下，斜上
下：row + 1  到达下边为止
斜上： row - 1, col + 1， 到达上边为止

*/

func convert(s string, numRows int) string {
	matrix, direction := make([][]byte, numRows, numRows), 0
	row, col := 0, 0

	for i, _ := range matrix {
		matrix[i] = make([]byte, len(s))
	}

	bytes := []byte(s)
	// z 填表
	for _, v := range bytes {
		matrix[row][col] = v
		if row == numRows-1 { // 到达下边界
			direction = 1
		}
		if row == 0 { // 到达上边界
			direction = 0
		}

		if direction == 0 {
			row++

		} else {
			row--
			col++
		}
	}

	// 输出
	ret := []byte{}
	for _, s := range matrix {
		for _, s1 := range s {

			if s1 != 0 {
				ret = append(ret, s1)
			}
		}

	}
	return string(ret)

}
