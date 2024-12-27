package leetcode

import (
	"fmt"
	"testing"

	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define

type question141 struct {
	para141
	ans141
}

// para 是参数
// one 代表第一个参数
type para141 struct {
	one []int
	pos int // 环的交点，-1 表示无环
}

// ans 是答案
// one 代表第一个答案
type ans141 struct {
	one bool
}

func Test_Problem141(t *testing.T) {

	qs := []question141{

		{
			para141{[]int{3, 2, 0, -4, 5, 7, 8, 4, 9}, 3},
			ans141{true},
		},

		{
			para141{[]int{3, 2, 0, -4, 5, 7, 8, 4, 9}, -1},
			ans141{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 141------------------------\n")

	for _, q := range qs {
		ans, p := q.ans141, q.para141
		if ans.one {
			// 表示有环
			fmt.Printf("[input]:%v\n[expect]:%v\n", p, ans)
			fmt.Printf("[output]:%v\n", detectCycle(structures.Ints2ListWithCycle(p.one, 3)))
		} else {
			// 无环
			fmt.Printf("[input]:%v\n[expect]:%v\n", p, ans)
			fmt.Printf("[output]:%v\n", detectCycle(structures.Ints2List(p.one)))
		}

	}
	fmt.Printf("\n\n\n")
}
