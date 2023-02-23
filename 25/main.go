package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(reverseKGroup(NewListNode(1, 2, 3, 4), 2))
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	var reverse func(head, foot *ListNode) *ListNode
	reverse = func(head, foot *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}

		newHead := reverse(head.Next, foot)
		head.Next.Next = head
		head.Next = nil

		return newHead
	}

	//var reverseRange func(head, foot *ListNode) *ListNode
	//reverseRange = func(head, foot *ListNode) *ListNode {
	//
	//}

	//count := 0
	//for cur, pre := head, new(ListNode); cur != nil; cur = cur.Next {
	//	count++
	//
	//	if count
	//}

	return reverse(head, head.Next.Next)
}
