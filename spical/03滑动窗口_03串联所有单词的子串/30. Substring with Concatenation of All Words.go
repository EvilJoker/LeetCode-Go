package leetcode

import (
	"maps"
)

/*
题目： 给一个单词数组，单词等长。 找出字符串包含 所有单词的起始索引。

思路：
1. 单词数组放入 map， 每次搞一个副本。
2. 从头开始遍历字符串字母，命中map 副本一次，就减一，直到全部命中次数归零，返回起始索引，
3. 因为单词等长，所以下一个单词是啥长度也确定

*/

func findSubstring(s string, words []string) []int {
	count := len(words)
	if count == 0 {
		return []int{}
	}

	word_len := len(words[0])
	wordslen := word_len * count
	if len(s) < wordslen {
		return []int{}
	}

	result := []int{}
	// 单词：次数
	dict := map[string]int{}
	for _, word := range words {
		dict[word] = dict[word] + 1
	}
	// 注意最大长度
	for i := 0; i < len(s)-wordslen+1; i++ {
		dict_copy := maps.Clone(dict)
		count_copy := count
		// 最后一个至少有一个单词
		for j := i; j < len(s)-word_len+1; {

			next_word := s[j : j+word_len]
			// 找不到,下一个
			if dict_copy[next_word] == 0 {
				break
			}
			// 找到,计数减一，右移
			dict_copy[next_word] = dict_copy[next_word] - 1
			count_copy--
			// 更新索引
			j += word_len
			// 全部找到
			if count_copy == 0 {
				result = append(result, i)
				break
			}
		}

	}
	return result

}

func findSubstring1(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := []int{}
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}
	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
	for i, start := 0, 0; i < len(s)-length+1 && start < len(s)-length+1; i++ {
		//fmt.Printf("sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
		if tmpCounter[s[i:i+length]] > 0 {
			tmpCounter[s[i:i+length]]--
			//fmt.Printf("******sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
			if checkWords(tmpCounter) && (i+length-start == totalLen) {
				res = append(res, start)
				continue
			}
			i = i + length - 1
		} else {
			start++
			i = start - 1
			tmpCounter = copyMap(counter)
		}
	}
	return res
}

func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}
