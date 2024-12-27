package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem125(t *testing.T) {

	tcs := []struct {
		s   string
		ans bool
	}{

		{
			"0p",
			false,
		},

		{
			"0",
			true,
		},

		{
			"race a car",
			false,
		},

		{
			"A man, a plan, a canal: Panama",
			true,
		},
	}
	fmt.Printf("------------------------Leetcode Problem 125------------------------\n")

	for _, tc := range tcs {
		fmt.Printf("[input]:%v       [expect]:%v\n", tc.s, tc.ans)
		fmt.Printf("[output]:%v\n", isPalindrome(tc.s))
	}
	fmt.Printf("\n\n\n")
}
