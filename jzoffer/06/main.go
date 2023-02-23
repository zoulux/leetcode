package main

import (
	"fmt"
	. "leetcode"
)

//输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

//示例 1：
//输入：head = [1,3,2]
//输出：[2,3,1]

func main() {
	fmt.Println(reversePrint(NewListNode(1, 3, 2)))
}

//func reversePrint(head *ListNode) []int {
//	var reverse func(*ListNode)
//	var values []int
//	reverse = func(node *ListNode) {
//		if node == nil {
//			return
//		}
//		reverse(node.Next)
//		values = append(values, node.Val)
//		return
//	}
//	reverse(head)
//
//	return values
//}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}
	return append(reversePrint(head.Next), head.Val)
}
