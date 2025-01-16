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
题目： 给出多个没有重叠空间（已经排序），然后把新的区间放进去能进行合并

思路：
想复杂了，笨办法：按照排序插入目标位置，然后从前向后合并
*/
func insert(intervals []Interval, newInterval Interval) []Interval {

	// 特殊情况
	if len(intervals) == 0 {
		return []Interval{newInterval}
	}
	//加入并排序
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	// 合并
	ret := []Interval{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := &ret[len(ret)-1]
		cur := intervals[i]
		// 不重叠
		if last.End < cur.Start {
			ret = append(ret, cur)
			continue
		}
		// 有重叠
		if last.End <= cur.End {
			// 不被包含，更新结果的值
			last.End = cur.End
		}
	} // 处理结尾的值

	return ret

}

func insertbk(intervals []Interval, newInterval Interval) []Interval {
	res := make([]Interval, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex].End < newInterval.Start {
		res = append(res, intervals[curIndex])
		curIndex++
	}

	for curIndex < len(intervals) && intervals[curIndex].Start <= newInterval.End {
		newInterval = Interval{Start: min(newInterval.Start, intervals[curIndex].Start), End: max(newInterval.End, intervals[curIndex].End)}
		curIndex++
	}
	res = append(res, newInterval)

	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
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
