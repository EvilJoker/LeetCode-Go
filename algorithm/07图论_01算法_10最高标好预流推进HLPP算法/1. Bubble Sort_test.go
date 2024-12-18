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
	nums []int
}

// ans 是答案
// one 代表第一个答案
type ans1 struct {
	one []int
}

func Test_Problem1(t *testing.T) {

	qs := []question1{

		{
			para1{[]int{0, 8, 7, 3, 3, 4, 2}},
			ans1{[]int{0, 2, 3, 3, 4, 7, 8}},
		},
		// 如需多个测试，可以复制上方元素。
	}

	fmt.Printf("------------------------Leetcode Problem 1------------------------\n")

	for _, q := range qs {
		ans, p := q.ans1, q.para1
		fmt.Printf("【input】:%v\n 【expect】:%v\n 【output】:%v\n", p, ans, sort(p.nums))
	}
	fmt.Printf("\n\n\n")
}
