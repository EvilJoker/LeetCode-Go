package leetcode

import (
	"fmt"
	"testing"
)

type question1742 struct {
	para1742
	ans1742
}

// para 是参数
// one 代表第一个参数
type para1742 struct {
	lowLimit  int
	highLimit int
}

// ans 是答案
// one 代表第一个答案
type ans1742 struct {
	one int
}

func Test_Problem1742(t *testing.T) {

	qs := []question1742{

		{
			para1742{1, 10},
			ans1742{2},
		},

		{
			para1742{5, 15},
			ans1742{2},
		},

		{
			para1742{19, 28},
			ans1742{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1742------------------------\n")

	for _, q := range qs {
		ans, p := q.ans1742, q.para1742
		fmt.Printf("[input]:%v\n[expect]:%v\n", p, ans)
		fmt.Printf("[output]:%v\n", countBalls(p.lowLimit, p.highLimit))
	}
	fmt.Printf("\n\n\n")
}
