package leetcode

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

/*
题目：
对字符串分组。不论排序，字符相同的单词为一组


思路：

对每个单词进行排序，如果字符串字符相同，那么经过排序的 key 就相同

*/

func groupAnagrams(strs []string) [][]string {
	// 中间数组，结果字符串
	record, res := map[string][]string{}, [][]string{}
	for _, str := range strs {
		sByte := []rune(str)

		sort.Sort(sortRunes(sByte))
		key := string(sByte)
		valueArr := record[key] // 排序后数组

		valueArr = append(valueArr, str)
		record[string(sByte)] = valueArr
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}
