package leetcode

import "strings"

/*
题目: Input: pattern = "abba", str = "dog cat cat dog", 判断是否同构

思路： 和同构字符串一样，区别是第二个字符串 是字符串数组.先把字符串分割，当数组用
*/
func wordPattern(pattern string, str string) bool {
	strArr := strings.Split(str, " ")
	// value 第一次出现的 index
	sMap := map[string]int{}
	pMap := map[byte]int{}

	if len(str) != len(strArr) {
		return false
	}

	for i := 0; i < len(str); i++ {
		p_index := pMap[pattern[i]]
		s_index := sMap[strArr[i]]

		// 索引不相同
		if p_index != s_index {
			return false
		}
		// 索引相同
		if p_index == 0 {
			pMap[pattern[i]] = i
			sMap[strArr[i]] = i
		}
	}
	return true
}
func wordPattern1(pattern string, str string) bool {
	strList := strings.Split(str, " ")
	patternByte := []byte(pattern)
	if pattern == "" || len(patternByte) != len(strList) {
		return false
	}

	pMap := map[byte]string{}
	sMap := map[string]byte{}
	for index, b := range patternByte {
		if _, ok := pMap[b]; !ok {
			if _, ok = sMap[strList[index]]; !ok {
				pMap[b] = strList[index]
				sMap[strList[index]] = b
			} else {
				if sMap[strList[index]] != b {
					return false
				}
			}
		} else {
			if pMap[b] != strList[index] {
				return false
			}
		}
	}
	return true
}
