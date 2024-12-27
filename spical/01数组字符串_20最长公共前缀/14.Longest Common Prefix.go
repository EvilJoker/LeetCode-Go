package leetcode

/*
	题目： 给定一个字符串数组，求所有字符串的公共前缀

	   Input: strs = ["flower","flow","flight"]

思路：

找到最小的字符串，开始遍历。当出现某个字母不是一个字符串的长度时。就终止
*/
func longestCommonPrefix(strs []string) string {
	prefix := ""
	index := 0
	for {
		char := ""

		for _, s := range strs {
			// 超出长度
			if index > len(s)-1 {
				return prefix
			}

			// 首次赋值
			if char == "" {
				char = string(s[index])
			}
			// 匹配
			if string(s[index]) != char {
				return prefix
			}

		}
		prefix += char
		index++

	}
	return prefix

}
