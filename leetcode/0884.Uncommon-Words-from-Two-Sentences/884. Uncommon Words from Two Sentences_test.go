package leetcode

import (
	"fmt"
	"testing"
)

type question884 struct {
	para884
	ans884
}

// para 是参数
// one 代表第一个参数
type para884 struct {
	A string
	B string
}

// ans 是答案
// one 代表第一个答案
type ans884 struct {
	one []string
}

func Test_Problem884(t *testing.T) {

	qs := []question884{

		{
			para884{"this apple is sweet", "this apple is sour"},
			ans884{[]string{"sweet", "sour"}},
		},

		{
			para884{"apple apple", "banana"},
			ans884{[]string{"banana"}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 884------------------------\n")

	for _, q := range qs {
		ans, p := q.ans884, q.para884
		fmt.Printf("[input]:%v\n[expect]:%v\n", p, ans)
		fmt.Printf("[output]:%v\n", uncommonFromSentences(p.A, p.B))
	}
	fmt.Printf("\n\n\n")
}
