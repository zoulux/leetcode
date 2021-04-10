package leetcode

import (
	. "mutil"
	"testing"
)

func getKthFromEnd(head *ListNode, k int) *ListNode {

	size := 0
	var helper func(head *ListNode) *ListNode
	helper = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		last := helper(head.Next)

		size++

		if size == k {
			return head
		} else if size < k {
			return nil
		} else {
			return last
		}
	}

	return helper(head)
}

func TestGetKthFromEnd(t *testing.T) {
	list := GenerateLists(1, 2, 3, 4, 5)
	t.Log(getKthFromEnd(list, 2))
}
