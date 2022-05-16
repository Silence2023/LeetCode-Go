package leetcode

import (
	"Leetcode-GO/structures"
)

// ListNode define
type ListNode = structures.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	node1, node2, carry, current := 0, 0, 0, head
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			node1 = 0
		} else {
			node1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			node2 = 0
		} else {
			node2 = l2.Val
			l2 = l2.Next
		}
		// 相加计算个位数
		current.Next = &ListNode{Val: (node1 + node2 + carry) % 10}
		current = current.Next
		// 获取相加后的十位数
		carry = (node1 + node2 + carry) / 10
	}
	return head.Next
}
