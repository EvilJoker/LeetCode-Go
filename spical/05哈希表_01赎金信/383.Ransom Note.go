package leetcode

/*
题目：magazine 字符串 是否可以 包含 ransomNote。 这个包含只是字符，不是指顺序
思路：设置map， 对 magazine 进行统计。
遍历ransomNote，如果从map 取不出来，就实名找不到
*/

func canConstruct(ransomNote string, magazine string) bool {

}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	var cnt [26]int
	// 26个字母
	for _, v := range magazine {
		cnt[v-'a']++
	}
	for _, v := range ransomNote {
		cnt[v-'a']--
		// 当取不出来时
		if cnt[v-'a'] < 0 {
			return false
		}
	}
	return true
}
