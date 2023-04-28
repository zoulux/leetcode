package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(pairSum(NewListNode(1, 2, 3, 4)))

}

func pairSum(head *ListNode) int {

	var reverse func(head *ListNode) *ListNode
	reverse = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		newHead := reverse(head.Next)
		head.Next.Next = head
		head.Next = nil
		return newHead
	}

	ln := 0
	cur := head
	for cur != nil {
		ln++
		cur = cur.Next
	}

	cur = head
	for i := 0; i < ln/2-1; i++ {
		cur = cur.Next
	}

	l2 := cur.Next
	cur.Next = nil
	l2 = reverse(l2)

	var ans int
	for i := 0; i < ln/2; i++ {
		if d := head.Val + l2.Val; d > ans {
			ans = d
		}

		head = head.Next
		l2 = l2.Next
	}
	return ans
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum2(head *ListNode) int {

	ls := make([]int, 0)
	for head != nil {
		ls = append(ls, head.Val)
		head = head.Next
	}

	var ans int
	for l, r := 0, len(ls)-1; l < r; l, r = l+1, r-1 {
		if d := ls[l] + ls[r]; d > ans {
			ans = d
		}
	}

	return ans
}
