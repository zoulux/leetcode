package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(oddEvenList(NewListNode(1, 2, 3, 4, 5)))
	//输入: head = [2,1,3,5,6,4,7]
	//	输出: [2,3,6,7,1,5,4]
	fmt.Println(oddEvenList(NewListNode(2, 1, 3, 5, 6, 4, 7)))
	fmt.Println(oddEvenList(NewListNode(1, 1)))

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var oddEvenList func(h1, h2 *ListNode) *ListNode
	oddEvenList = func(h1, h2 *ListNode) *ListNode {
		if h1 == nil && h2 == nil {
			// 奇数链和偶数链都为空结束
			return nil
		}
		if h1 == nil {
			// 奇数链为空，则处理偶数链
			if h2.Next != nil {
				h2.Next = oddEvenList(h1, h2.Next.Next)
			}
			return h2
		}

		if h1.Next == nil {
			// 表示是奇数链的最后一个节点了
			h1.Next = oddEvenList(h1.Next, h2)
			return h1
		}

		h1.Next = oddEvenList(h1.Next.Next, h2)
		return h1
	}

	// h1 表示奇数链
	// h2 表示偶数链
	return oddEvenList(head, head.Next)
}
