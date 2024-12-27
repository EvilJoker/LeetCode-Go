package leetcode

/*
	题目：整数转罗马数字

思路： 刚好相反

解题思路采用贪心算法。将 1-3999 范围内的罗马数字从大到小放在数组中，从头选择到尾，即可把整数转成罗马数字。
*/
var symbols = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
var values = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	// 尽可能取大数
	ret := ""

	for i, v := range values {
		// 循环减去最大值
		for num >= v {
			ret += symbols[i]
			num -= v
			// fmt.Println(num)
		}
	}

	return ret

}
