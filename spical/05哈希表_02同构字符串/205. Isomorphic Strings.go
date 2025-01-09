package leetcode

/*
题目：保证两个字符串的结构一样, 比如  和 title 是一样的。
p 1，3 都是第一次出现

思路：

同时遍历生成两个字符串，生成 map. map 中的value 保存第一次出现的索引。

接口相同时，每次遍历两个字符串取的索引应该是相同的


*/

func isIsomorphic(s string, t string) bool {
	sMap, tMap := map[byte]int{}, map[byte]int{}

	for i := 0; i < len(s); i++ {
		s_index := sMap[s[i]]
		t_index := tMap[t[i]]

		// 不相等
		if s_index != t_index {
			return false
		}

		// 此时，索引相等，判断一个就可以。赋值
		if s_index == 0 {
			sMap[s[i]] = i
			tMap[t[i]] = i

		}

	}

	return true

}

func isIsomorphic1(s string, t string) bool {
	strList := []byte(t)
	patternByte := []byte(s)
	if (s == "" && t != "") || (len(patternByte) != len(strList)) {
		return false
	}

	pMap := map[byte]byte{}
	sMap := map[byte]byte{}
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
