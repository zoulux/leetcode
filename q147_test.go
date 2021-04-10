package leetcode

import (
	. "mutil"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	// 4 ,2,1,3
	// 1 ,4,2,3

	// sortLast := head

	// for sortLast != nil {

	// 	for cur != sortLast.Next {

	// 		// cur.Val<

	// 	}

	// }

	return nil

}

func TestInsertionSortList(t *testing.T) {
	l := GenerateLists(4, 2, 1, 3)
	t.Log(insertionSortList(l))
}
