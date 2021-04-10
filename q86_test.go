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
func partition4(root *ListNode, x int) *ListNode {

	small := new(ListNode)
	big := new(ListNode)

	cur := root
	curBig := big
	curSmall := small

	for cur != nil {
		if cur.Val >= x {
			curBig.Next = cur
			curBig = curBig.Next
		} else {
			curSmall.Next = cur
			curSmall = curSmall.Next

		}
		cur = cur.Next
	}

	curBig.Next = nil // 这个很重要 否则会形成环
	curSmall.Next = big.Next

	return small.Next
}

func TestPartition4(t *testing.T) {
	list := GenerateLists(1, 4, 3, 2, 5, 2)
	t.Log(partition4(list, 3))

}
