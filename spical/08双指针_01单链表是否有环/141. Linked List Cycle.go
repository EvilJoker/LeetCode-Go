package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
思路： 快慢指针
快指针 2倍速跑，跟操场跑步一样，有环就会套圈.
快的要保证 自己 和next 都不为 nill 否则会报错
*/
func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil { // 快判断2个
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}

	}

	return false
}
