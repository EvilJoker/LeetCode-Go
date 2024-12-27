package leetcode

import (
	"maps"
)

/*
	题目

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
返回一个最小的连续窗口，包含 T 的所有 字符(假如 存在多个B , 那么窗口也需要有多个 B)

思路：
1. 创建 T 的位图，一个 map
2. 使用双指针不断更新位图，匹配是够满足 T。
*/
func minWindow(s string, t string) string {

	left, right := 0, 0
	min_left, min_right, min_len := 0, 0, len(s)+1
	println(len(s))

	tdict := map[byte]int{}
	for right < len(s) {
		// 右移右边界，尝试增加
		c := s[right]
		tdict[c] = tdict[c] + 1

		for left <= right && maps_container(tdict, t) {
			// fmt.Printf("left %d, right%d sub %v, dict %v\n", left, right, s[left:right+1], tdict)
			if right-left+1 < min_len {
				min_len = right - left + 1
				min_left = left
				min_right = right
			}

			//包含时， 持续右移左边界，尝试缩减
			c := s[left]
			tdict[c] = tdict[c] - 1
			left++
		}

		right++

	}

	if min_len > len(s) {
		// 没找到
		return ""
	}
	// 右边界是不包含的
	return s[min_left : min_right+1]

}

// dicts 是否包含 string
func maps_container(dict_cop map[byte]int, t string) bool {
	dict := maps.Clone(dict_cop)
	for i := 0; i < len(t); i++ {
		c := t[i]
		if dict[c] == 0 {
			// 不存在
			return false
		}
		dict[c] = dict[c] - 1

	}
	return true

}

func minWindow1(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}
	var tFreq, sFreq [256]int
	result, left, right, finalLeft, finalRight, minW, count := "", 0, -1, -1, -1, len(s)+1, 0

	for i := 0; i < len(t); i++ {
		tFreq[t[i]-'a']++
	}

	for left < len(s) {
		if right+1 < len(s) && count < len(t) {
			sFreq[s[right+1]-'a']++
			if sFreq[s[right+1]-'a'] <= tFreq[s[right+1]-'a'] {
				count++
			}
			right++
		} else {
			if right-left+1 < minW && count == len(t) {
				minW = right - left + 1
				finalLeft = left
				finalRight = right
			}
			if sFreq[s[left]-'a'] == tFreq[s[left]-'a'] {
				count--
			}
			sFreq[s[left]-'a']--
			left++
		}
	}
	if finalLeft != -1 {
		result = string(s[finalLeft : finalRight+1])
	}
	return result
}
