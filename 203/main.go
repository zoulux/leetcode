package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(removeElements(NewListNode(1, 2, 3, 4), 4))
	fmt.Println(removeElements(NewListNode(1, 2, 6, 3, 4, 5, 6), 6))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElements(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head

}
func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	// 1,2,3,4
	// 1
	dummy := new(ListNode)
	dummy.Next = head

	cur := dummy.Next
	pre := dummy
	for cur != nil {
		if cur.Val == val {
			pre.Next = pre.Next.Next
			cur = pre.Next
		} else {
			cur = cur.Next
			pre = pre.Next
		}
	}

	return dummy.Next
}
