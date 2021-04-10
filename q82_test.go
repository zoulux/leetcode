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
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {
		for head != nil && head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		return deleteDuplicates2(head.Next)

	} else {
		head.Next = deleteDuplicates2(head.Next)
	}

	return head
}

func TestDeleteDuplicates2(t *testing.T) {
	list := GenerateLists(1, 2, 3, 3, 4, 4, 5)
	// list := GenerateLists(1, 1, 1, 2, 3)

	t.Log(deleteDuplicates2(list))

}
