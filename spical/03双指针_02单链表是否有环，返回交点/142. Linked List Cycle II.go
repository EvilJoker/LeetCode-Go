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
先判断是否有环：
快指针 2倍速跑，跟操场跑步一样，有环就会套圈.
快的要保证 自己 和next 都不为 nill

有环时（fast slow 相等时）--- 此时为相遇位置 q

将 slow 再次指向 head, 快慢指针单倍速跑，再次相遇就是 起点 p

// 原理 ： head 到环起点，距离为 l1, 起点到相遇位置为 l2
1. slow 走过的距离 ，l1+ l2
2. fast 走过距离， l1 + l2+ R（多一圈）
3. 因为 fast 2倍速 l1 + l2+ R = 2（l1 +l2） ---> l1 = R - l2 ---> R-l2 的含义是，从相遇点走到起点的距离
4. 也就是说，再走l1 距离就是起点，所以让另一个指针从头开始


*/

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow { // 有环
			fast = head // 重置到起点

			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast //相遇出循环

		}

	}
	return nil
}
