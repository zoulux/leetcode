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
func partition5(head *ListNode, x int) *ListNode {

	small := new(ListNode)
	larger := new(ListNode)
	cur, smallCur, largerCur := head, small, larger

	for cur != nil {
		if cur.Val > x {
			largerCur.Next = cur
			largerCur = largerCur.Next

		} else {
			smallCur.Next = cur
			smallCur = smallCur.Next

		}
		cur = cur.Next
	}
	largerCur.Next = nil
	smallCur.Next = larger.Next

	return small.Next

}

func TestPartition5(t *testing.T) {
	partition5(nil, 0)
}
