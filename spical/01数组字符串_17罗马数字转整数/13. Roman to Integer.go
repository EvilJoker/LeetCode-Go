package leetcode

import "fmt"

/*
题目： 翻译罗马数字为 10进制
1 I
2 II
3 III
4 IV
5 V
6 VI
...
10 X
11 XI
...
14 XIV
...
31 XXXI
40 XL

注意正常数字是 IVXLCDM 如果出现一次逆序，就是要减的意思

思路：
使用map 创建基础的对照表。
遍历数组，判断 I X C 的 右边有没有 VXLCDM ,如果符合情况就当成一个数字

优化： 正常情况加遇到右边数字是加，但是如果遇到 大一位的就变成减

*/

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// MCMXICIVI
// M 1000  CM 900 X 10 I 1 C 100 IV 4 I 1 --2016
// M 1000 CM 900 X 10 IC 99 IV 4 I 1 -- 2014 -- 正确
func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	sum := 0

	lastNum := roman[string(s[0])]
	for i := 1; i < len(s); i++ {
		// 如果当前数比上一个数大，就是说明逆序，要减
		curNum := roman[string(s[i])]

		if curNum > lastNum {
			sum -= lastNum
		} else {
			sum += lastNum
		}

		lastNum = curNum

	}
	fmt.Println(lastNum)

	return sum + lastNum

}
