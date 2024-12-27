package leetcode

import (
	"strings"
)

/*
思路：回文串特点，从两头遍历，他的字符都相等。
特殊点，这个题需要忽略掉符号，空格，大小写
*/

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	bytes := []byte(s)

	start, end := 0, len(bytes)-1

	for start < end {
		// 忽略字符
		if !isChar(bytes[start]) {
			start++
			continue
		}
		if !isChar(bytes[end]) {
			end--
			continue
		}

		// 回文匹配
		if bytes[start] != bytes[end] {
			return false
		}
		start++
		end--

	}

	return true

}

// 判断是否是数字或者字符串
func isChar(c byte) bool {
	if ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') {
		return true
	}

	return false
}
