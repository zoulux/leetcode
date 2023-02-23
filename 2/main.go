package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	//fmt.Println(NewListNode(2, 4, 3))
	fmt.Println(addTwoNumbers(NewListNode(2, 4, 3), NewListNode(5, 6, 4)))
	fmt.Println(addTwoNumbers(NewListNode(9, 9, 9, 9, 9, 9, 9), NewListNode(9, 9, 9, 9)))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var add func(l1, l2 *ListNode, x int) *ListNode
	add = func(l1, l2 *ListNode, x int) *ListNode {
		if l1 == nil && l2 == nil && x == 0 {
			return nil
		}

		if l1 != nil {
			x += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			x += l2.Val
			l2 = l2.Next
		}

		n := &ListNode{
			Val: x % 10,
		}
		n.Next = add(l1, l2, x/10)
		return n
	}
	return add(l1, l2, 0)
}
