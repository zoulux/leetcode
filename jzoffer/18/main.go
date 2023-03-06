package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(deleteNode(NewListNode(4, 5, 1, 9), 4))
	fmt.Println(deleteNode(NewListNode(4, 5, 1, 9), 5))
	fmt.Println(deleteNode(NewListNode(4, 5, 1, 9), 1))
	// 4 -> 5 -> 1 -> 9
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}

	head.Next = deleteNode(head.Next, val)
	return head
}
