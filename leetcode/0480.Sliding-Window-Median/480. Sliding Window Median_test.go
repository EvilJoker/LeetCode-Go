package leetcode

import (
	"fmt"
	"testing"
)

type question480 struct {
	para480
	ans480
}

// para 是参数
// one 代表第一个参数
type para480 struct {
	one []int
	k   int
}

// ans 是答案
// one 代表第一个答案
type ans480 struct {
	one []int
}

func Test_Problem480(t *testing.T) {

	qs := []question480{

		{
			para480{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3},
			ans480{[]int{1, -1, -1, 3, 5, 6}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 480------------------------\n")

	for _, q := range qs {
		ans, p := q.ans480, q.para480
		fmt.Printf("[input]:%v\n[expect]:%v\n", p, ans)
		fmt.Printf("[output]:%v\n", medianSlidingWindow1(p.one, p.k))
	}
	fmt.Printf("\n\n\n")
}
