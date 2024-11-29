package leetcode

import (
	"fmt"
	"testing"
)

type question810 struct {
	para810
	ans810
}

// para 是参数
// one 代表第一个参数
type para810 struct {
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans810 struct {
	one bool
}

func Test_Problem810(t *testing.T) {

	qs := []question810{

		{
			para810{[]int{1, 1, 2}},
			ans810{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 810------------------------\n")

	for _, q := range qs {
		ans, p := q.ans810, q.para810
		fmt.Printf("【input】:%v\n 【expect】:%v\n 【output】:%v\n", p, ans, xorGame(p.nums))
	}
	fmt.Printf("\n\n\n")
}
