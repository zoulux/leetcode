package main

import (
	"fmt"
	. "leetcode"
)

func main() {
	fmt.Println(getKthFromEnd2(NewListNode(1, 2, 3, 4, 5), 1))
	fmt.Println(getKthFromEnd2(NewListNode(1, 2, 3, 4, 5), 2))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {

	var travel func(head *ListNode) (int, *ListNode)
	travel = func(head *ListNode) (int, *ListNode) {
		if head == nil {
			return 0, nil
		}

		x, h := travel(head.Next)
		x++
		if x == k {
			return k, head
		}
		return x, h
	}

	_, t := travel(head)
	return t
}

func getKthFromEnd2(head *ListNode, k int) *ListNode {
	var travel func(head *ListNode) *ListNode
	travel = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}

		node := travel(head.Next)
		if k == 0 {
			return node
		}
		k--
		return head
	}

	return travel(head)
}
