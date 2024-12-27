package leetcode

/*
	子序列：判断s 是否是 t 的子串

使用两个指针，相同就更新索引，直到 全部匹配
*/
func isSubsequence(s string, t string) bool {
	m, n := len(s), len(t)

	// s 的索引
	p := 0

	for i := 0; i < n && p < m; i++ {
		// 相同就更新 p
		if t[i] == s[p] {
			p++
		}
	}

	// 全部匹配
	if p >= m-1 {
		return true
	}

	return false

}
