package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(mergeTwoLists(NewListNode(1, 2, 4), NewListNode(1, 3, 4)))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1.Next = mergeTwoLists(l1.Next, l2)
			return l1
		} else {
			l2.Next = mergeTwoLists(l1, l2.Next)
			return l2
		}
	}

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	return nil
}
