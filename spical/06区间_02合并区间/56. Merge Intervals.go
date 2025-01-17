package leetcode

import (
	"sort"

	"github.com/halfrost/LeetCode-Go/structures"
)

// Interval define
type Interval = structures.Interval

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

/*
题目： 合并无序区间，比如 [1,3] [2,6] -> [1,6]

思路： 按首数字，对区间从小到大进行排序
遍历： 如果发现当前和上一个区间有重叠就进行合并
*/
func merge56(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	ret := []Interval{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	ret = append(ret, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := &ret[len(ret)-1] // 取地址，而不是copy
		// 不合并
		if last.End < intervals[i].Start {
			ret = append(ret, intervals[i])
			continue
		}
		// 合并

		last.End = intervals[i].End
	}
	return ret

}

func merge56bk(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	quickSort(intervals, 0, len(intervals)-1)
	res := make([]Interval, 0)
	res = append(res, intervals[0])
	curIndex := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start > res[curIndex].End {
			curIndex++
			res = append(res, intervals[i])
		} else {
			res[curIndex].End = max(intervals[i].End, res[curIndex].End)
		}
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func partitionSort(a []Interval, lo, hi int) int {
	pivot := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if (a[j].Start < pivot.Start) || (a[j].Start == pivot.Start && a[j].End < pivot.End) {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}
func quickSort(a []Interval, lo, hi int) {
	if lo >= hi {
		return
	}
	p := partitionSort(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}
