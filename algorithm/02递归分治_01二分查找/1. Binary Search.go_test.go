package leetcode

import (
	"fmt"
	"testing"
)

type question1 struct {
	para1
	ans1
}

// para 是参数
// one 代表第一个参数
type para1 struct {
	nums   []int
	target int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{

		{
			para1{[]int{1, 3, 5, 7, 9, 11, 13, 15, 17}, 7},
			ans1{3},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		ans, p := q.ans1, q.para1
		fmt.Printf("[input]:%v\n[expect]:%v\n", p, ans)
		fmt.Printf("[output]:%v\n", binarySearch(p.nums, p.target))
	}
	fmt.Printf("\n\n\n")
}
