package structures

import (
	"fmt"
)

// ListNode 是链接节点
// 这个不能复制到*_test.go文件中。会导致Travis失败
type ListNode struct {
	Val  int
	Next *ListNode
}

// List2Ints convert List to []int
func List2Ints(head *ListNode) []int {
	// 链条深度限制，链条深度超出此限制，会 panic
	limit := 100

	times := 0

	res := []int{}
	for head != nil {
		times++
		if times > limit {
			msg := fmt.Sprintf("链条深度超过%d，可能出现环状链条。请检查错误，或者放宽 l2s 函数中 limit 的限制。", limit)
			panic(msg)
		}

		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// Ints2List convert []int to List
func Ints2List(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

// GetNodeWith returns the first node with val
func (l *ListNode) GetNodeWith(val int) *ListNode {
	res := l
	for res != nil {
		if res.Val == val {
			break
		}
		res = res.Next
	}
	return res
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle // 构造有环的 列表，pos 是环的尾巴
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	// 得到环的交点
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	// 找到单链表的尾巴节点
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	// 尾巴节点指向头结点得到环
	tail.Next = c
	return head
}
