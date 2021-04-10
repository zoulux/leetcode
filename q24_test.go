package leetcode

import (
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	next := head.Next

	head.Next = swapPairs(next.Next)
	next.Next = head

	return next
}

func TestSwapPairs(t *testing.T) {
	swapPairs(GenerateLists(1, 2, 3, 4))
}
