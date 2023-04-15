package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(deleteDuplicates(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{Val: 5},
				},
			},
		},
	}))
}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	return head
}

//leetcode submit region end(Prohibit modification and deletion)
