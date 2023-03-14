package main

import . "leetcode"

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {

	fast, slow := head, head
	var enter *ListNode
	for slow != nil && fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next

		if fast == slow {
			enter = fast
			break
		}
	}

	if enter == nil {
		// 无环
		return nil
	}

	slow = head

	for slow != enter {
		slow = slow.Next
		enter = enter.Next
	}

	return slow
}
