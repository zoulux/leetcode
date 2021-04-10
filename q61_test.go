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
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	len := getLen(head)
	k = k % len // 只需要旋转 k 次

	if k == 0 {
		return head
	}

	first, seccond := head, head

	for i := 0; i < k; i++ {
		first = first.Next
	}

	for first.Next != nil {
		first = first.Next
		seccond = seccond.Next
	}

	seccond, seccond.Next = seccond.Next, nil

	cur := seccond
	if cur != nil {
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = head
	}

	return seccond
}

func TestRotateRight(t *testing.T) {
	// t.Log(rotateRight(GenerateLists(1, 2, 3, 4, 5), 2))
	// t.Log(rotateRight(GenerateLists(1, 2), 2))
	t.Log(rotateRight(GenerateLists(1, 2), 1))
	t.Log(rotateRight(&ListNode{Val: 1}, 1))

}

func getLen(head *ListNode) int {
	cur := head

	length := 0
	for cur != nil {
		length++
		cur = cur.Next
	}

	return length
}
