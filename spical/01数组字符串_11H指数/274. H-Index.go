package leetcode

import (
	"slices"
)

/*
题目： 求 h-index。h-index 值的定义：如果他/她的 N 篇论文中至少有 h 引用，而其他 N-h 论文的引用数不超过 h 引用数。


[3,0,6,1,5]  -->3
数组代表每篇论文的引用数， 其中3篇 >3, 其他两篇 不大于 3， 那么3 就是引用数

解析：
假设有 n 篇论文，每篇有自己的引用数。

a 篇论文引用 > h
n - a 篇数量 < h





思路：
将数组升序排序，找到满足上数公式的值。上面的公式没有实际意义，不好理解

a 篇论文引用 >= h  ---> n-
a - h 数量 =< h

*/
// 解法一
func hIndex(citations []int) int {
	// 排序
	slices.Sort(citations)
	n := len(citations)

	for i, v := range citations {
		h := n - i // i篇论文 <= citations[i], 剩余的 n -i 篇 >= citations[i]

		// 只要 v>=h 所有条件都满足了。就说明 右侧有 h 篇
		if v >= h {
			return h
		}

	}

	return 0

}
