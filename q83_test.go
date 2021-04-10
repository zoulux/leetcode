package leetcode

import (
	. "mutil"
	"testing"
)

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Val == head.Next.Val {

		head.Next = head.Next.Next // 当前值和下个值相等，指向下下个

		deleteDuplicates(head) // 这里是个坑 因为当前值可能又和下个值相等了，所以继续从当前节点递归

	} else {
		deleteDuplicates(head.Next) //不相等从下个值继续递归
	}
	return head
}

func TestDeleteDuplicates(t *testing.T) {
	head := GenerateLists(1, 1, 2)
	t.Log(head)
}
