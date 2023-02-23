package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(mergeTwoLists(NewListNode(1, 2, 4), NewListNode(1, 3, 4)))

}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list1.Next = mergeTwoLists(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists(list1, list2.Next)
			return list2
		}
	}

	if list1 != nil {
		return list1
	}

	if list2 != nil {
		return list2
	}

	return nil
}
