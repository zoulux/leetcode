package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	//输入：l1 = [7,2,4,3], l2 = [5,6,4] 输出：[7,8,0,7]
	fmt.Println(addTwoNumbers(NewListNode(7, 2, 4, 3), NewListNode(5, 6, 4)))
	fmt.Println(addTwoNumbers(NewListNode(7), NewListNode(3)))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var revert func(n *ListNode) *ListNode
	revert = func(n *ListNode) *ListNode {
		if n == nil || n.Next == nil {
			return n
		}

		h := revert(n.Next)
		n.Next.Next = n
		n.Next = nil
		return h
	}

	lt1 := revert(l1)
	lt2 := revert(l2)

	var add func(l1, l2 *ListNode, carry int) *ListNode
	add = func(l1, l2 *ListNode, carry int) *ListNode {
		if l1 == nil && l2 == nil && carry == 0 {
			return nil
		}

		node := &ListNode{}
		if l1 != nil {
			carry += l1.Val
		}
		if l2 != nil {
			carry += l2.Val
		}

		node.Val = carry % 10

		if l1 != nil && l2 != nil {
			node.Next = add(l1.Next, l2.Next, carry/10)
		} else if l1 == nil && l2 != nil {
			node.Next = add(nil, l2.Next, carry/10)
		} else if l1 != nil && l2 == nil {
			node.Next = add(l1.Next, nil, carry/10)
		}
		return node
	}
	return revert(add(lt1, lt2, 0))
}
