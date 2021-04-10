package leetcode

import (
	"testing"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val: 0,
	}

	carry := 0
	current := head

	for l1 != nil || l2 != nil || carry != 0 {

		v := carry

		if l1 != nil {
			v += l1.Val
			l1 = l1.Next

		}

		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val: v % 10,
		}

		current = current.Next

		carry = v / 10
	}

	return head.Next
}

func TestAddTwoNumbers(t *testing.T) {
	// l1 = [2,4,3], l2 = [5,6,4]
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	ret := addTwoNumbers(l1, l2)

	for ret != nil {
		t.Logf("%v", ret.Val)
		ret = ret.Next
	}

}
