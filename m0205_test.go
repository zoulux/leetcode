package leetcode

import (
	. "mutil"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {

	tail := 0
	dummy := new(ListNode)
	var helper func(root, l1, l2 *ListNode)

	helper = func(root, l1, l2 *ListNode) {

		if l1 == nil || l2 == nil {

			return
		}

		ret := l1.Val + l2.Val + tail
		tail = ret / 10
		ret = ret % 10

		helper(root, l1.Next, l2.Next)

		dummy.Val = last.Val
		dummy = dummy.Next

	}
	helper(new(ListNode), l1, l2)

	return dummy.Next
}

func TestAddTwoNumbers22(t *testing.T) {
	t.Log(addTwoNumbers2(GenerateLists(7, 1, 6), GenerateLists(5, 9, 2)))
}
