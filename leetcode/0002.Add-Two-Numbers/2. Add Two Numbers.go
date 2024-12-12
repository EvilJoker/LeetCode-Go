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

/* 思路
```
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
```
Explanation: 342 + 465 = 807.

链表从地位开始相加，相加>10, 余数作为新节点，进位为1。下一个数是进位加两个节点，直到为空
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	pre := result
	up := 0 // 进位

	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += up

		up = sum / 10
		pre.Next = &ListNode{Val: sum % 10}
		pre = pre.Next

	}
	// 如果最后还有进位，则需要在最后加一个节点
	if up > 0 {
		pre.Next = &ListNode{Val: up}
	}

	return result.Next
}
