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
			// 这里可以优化一下，当l1为空，并且carry为0时，意味着l2的下一节点不会再改变了。
			if carry != 0 {
				node1 = 0
			} else if l2 != nil {
				current.Next = l2
				break
			}
		} else {
			node1 = l1.Val
		}
		if l2 == nil {
			if carry != 0 {
				node2 = 0
				if l1 != nil {
					l1 = l1.Next
				}
			} else if l1 != nil {
				current.Next = l1
				break
			}
		} else {
			node2 = l2.Val
			l2 = l2.Next
			if l1 != nil {
				l1 = l1.Next
			}
		}
		// 相加计算个位数
		current.Next = &ListNode{Val: (node1 + node2 + carry) % 10}
		current = current.Next
		// 获取相加后的十位数
		carry = (node1 + node2 + carry) / 10
	}
	return head.Next
}
