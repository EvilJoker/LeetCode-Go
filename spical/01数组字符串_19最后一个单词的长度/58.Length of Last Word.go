package leetcode

/*
题目：求最后一个单词的长度

思路：倒序遍历字符串，第一个出现的空格为分界点。注意单词前后有空格
*/
func lengthOfLastWord(s string) int {
	n := len(s)
	// 找第一个非空格索引,忽略空格
	start := n - 1
	for s[start] == ' ' {
		start--
	}

	// 找随后的空格，忽略非空格
	end := start
	for s[end] != ' ' {
		end--
	}

	return start - end

}
