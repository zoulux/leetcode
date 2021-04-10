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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)

	dummy.Next = head

	curiPre := dummy
	curi := curiPre.Next
	for i := 1; i < left; i++ {
		curiPre, curi = curi, curi.Next
	}

	curj := curi
	for j := left; j < right; j++ {
		curj = curj.Next
	}

	curjNext := curj.Next
	curj.Next = nil

	curiPre.Next = reverse22(curi)
	curi.Next = curjNext

	return dummy.Next
}

func reverse22(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse22(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

var last *ListNode
var first *ListNode

func reverse23(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverse23(head, left-1, right-1)

	return head
}

var successor *ListNode

func reverseN(head *ListNode, m int) *ListNode {

	if m == 1 {
		successor = head.Next
		return head
	}

	last := reverseN(head.Next, m-1)

	head.Next.Next = head
	head.Next = successor

	return last
}

func TestReverseBetween(t *testing.T) {
	list := GenerateLists(1, 2, 3, 4, 5)

	t.Log(reverse23(list, 2, 4))

}
