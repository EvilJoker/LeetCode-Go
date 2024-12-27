package leetcode

/*
	题目 翻转单词

[input]:{the sky is blue}

[expect]:{blue is sky the}

思路：
翻转两次，先整体翻转，然后每个单词再翻转一次
*/
func reverseWords151(s string) string {
	bytes := []byte(s)
	n := len(s)
	// 整体翻转
	reverseSingle12(bytes)

	// 按照单词翻转
	start := 0
	end := 0
	for end < n-1 {
		// 忽略空格
		for {
			if start < n && bytes[start] == ' ' {
				start++
				continue
			}
			break
		}
		// 忽略非空格
		end = start
		for {
			if end < n && bytes[end] != ' ' {
				end++
				continue
			}
			break
		}

		reverseSingle12(bytes[start:end])
		start = end

	}

	return string(bytes)

}

func reverseSingle12(bytes []byte) {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}
}
