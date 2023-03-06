package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(getIntersectionNode(NewListNode(1, 2), NewListNode(1, 2, 3)))

}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var travel func(A, B *ListNode) *ListNode
	travel = func(A, B *ListNode) *ListNode {

		if A == B {
			return A
		}

		if A != nil {
			A = A.Next
		} else {
			A = headB
		}

		if B != nil {
			B = B.Next
		} else {
			B = headA
		}
		return travel(A, B)
	}
	return travel(headA, headB)
}
