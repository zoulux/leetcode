package main

import (
	"fmt"
	. "leetcode"
)

//定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

// 示例:
// 输入: 1->2->3->4->5->NULL
// 输出: 5->4->3->2->1->NULL

func main() {
	fmt.Println(reverseList(NewListNode(1, 2, 3, 4, 5)))
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	var reverse func(cur, pre *ListNode) *ListNode
	reverse = func(cur, pre *ListNode) *ListNode {
		if cur == nil {
			return pre
		}

		tmp := cur.Next
		cur.Next = pre
		return reverse(tmp, cur)
	}
	return reverse(head, nil)
}
