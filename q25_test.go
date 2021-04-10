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
func reverseKGroup(head *ListNode, k int) *ListNode {

	b := head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(head, b)

	head.Next = reverseKGroup(b, k)

	return newHead
}

/// [a, b) 左闭右开
func reverse(a, b *ListNode) (pre *ListNode) {

	for cur := a; cur != b; {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return
}

func TestReverseKGroup(t *testing.T) {
	list := GenerateLists(1, 2, 3, 4, 5, 6)
	// t.Log(reverse(list, list.Next.Next))
	t.Log(reverseKGroup(list, 2))

}
