package leetcode

import (
	"fmt"
	"testing"
)

type question134 struct {
	para134
	ans134
}

// para 是参数
// one 代表第一个参数
type para134 struct {
	one []int
	two []int
}

// ans 是答案
// one 代表第一个答案
type ans134 struct {
	one int
}

func Test_Problem134(t *testing.T) {
	qs := []question134{

		{
			para134{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
			ans134{3},
		},

		{
			para134{[]int{2, 3, 4}, []int{3, 4, 3}},
			ans134{-1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 134------------------------\n")

	for _, q := range qs {
		ans, p := q.ans134, q.para134
		fmt.Printf("[input]:%v, %v       [expect]:%v\n", p.one, p.two, ans.one)
		fmt.Printf("[output]:%v\n", canCompleteCircuit(p.one, p.two))
		fmt.Printf("\n\n\n")
	}

}
