package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	d := NewListNode(1, 3, 4, 7, 1, 2, 6)
	fmt.Println(deleteMiddle(d))

	d = NewListNode(1, 2, 3, 4)
	fmt.Println(deleteMiddle(d))

	d = NewListNode(2, 1)
	fmt.Println(deleteMiddle(d))

	d = NewListNode(1)
	fmt.Println(deleteMiddle(d))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	dummy := new(ListNode)
	dummy.Next = head
	first, second := dummy, dummy.Next.Next
	//1,3,4,1,2,6

	for second != nil && second.Next != nil {
		first = first.Next
		second = second.Next
	}

	first.Next = first.Next.Next
	return head
}
