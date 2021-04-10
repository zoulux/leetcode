package leetcode

import (
	. "mutil"
	"testing"
)

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && slow != nil {
		if slow.Next != nil {
			slow = slow.Next
		} else {
			return false // 遇到空节点则无环
		}

		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
		} else {
			return false // 遇到空节点则无环
		}

		if slow == fast {
			return true // 相遇则有环
		}
	}
	return false
}

func TestHasCycle(t *testing.T) {

}
