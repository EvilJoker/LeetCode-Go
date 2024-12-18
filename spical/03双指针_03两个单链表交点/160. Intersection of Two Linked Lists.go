package leetcode

import (
	"fmt"

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
思路：两个链表的交点。
双指针，两个链表有交点，意味着交点到结尾这一段是相同的。但是起点到交点各不相同

方法1： 笨办法遍历计算长度
方法2： 拼接的思路。
a:------- + b:--- = b:--- + a:--------
当 a b 同时出发，第一个链表时，因为长度不同无法相遇，但是继续走时，发现第一次遍历消灭了差距，再此相遇就是交点



*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}
	p1, p2 := headA, headB

	for p1 != nil && p2 != nil && p1 != p2 {
		if p1 == p2 {
			return p1
		}

	}

	return nil

}
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	//boundary check
	if headA == nil || headB == nil {
		return nil
	}

	a := headA
	b := headB

	//if a & b have different len, then we will stop the loop after second iteration
	for a != b {
		//for the end of first iteration, we just reset the pointer to the head of another linkedlist
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
		fmt.Printf("a = %v b = %v\n", a, b)
	}
	return a
}
