package leetcode

import (
	"fmt"
	"testing"
)

type question520 struct {
	para520
	ans520
}

// para 是参数
type para520 struct {
	word string
}

// ans 是答案
type ans520 struct {
	ans bool
}

func Test_Problem520(t *testing.T) {

	qs := []question520{

		{
			para520{"USA"},
			ans520{true},
		},

		{
			para520{"FlaG"},
			ans520{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 520------------------------\n")

	for _, q := range qs {
		ans, p := q.ans520, q.para520
		fmt.Printf("【input】:%v\n 【expect】:%v\n 【output】:%v\n", p, ans, detectCapitalUse(p.word))
	}
	fmt.Printf("\n\n\n")
}
