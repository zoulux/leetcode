package main

import (
	"fmt"
	. "leetcode"
	"sort"
)

func main() {

	fmt.Println(sortList(NewListNode(1, 3, 2, 7, 4)))
}
func sortList(head *ListNode) *ListNode {

	merge := func(left, right *ListNode) *ListNode {
		dummy := new(ListNode)

		cur := dummy

		for left != nil && right != nil {
			if left.Val < right.Val {
				cur.Next = left
				left = left.Next
			} else {
				cur.Next = right
				right = right.Next
			}
			cur = cur.Next
		}
		if left != nil {
			cur.Next = left
		}

		if right != nil {
			cur.Next = right
		}

		return dummy.Next
	}

	var mergeSort func(head, tail *ListNode) *ListNode
	mergeSort = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return head
		}

		if head.Next == tail {
			head.Next = nil
			return head
		}

		slow, fast := head, head
		for fast != tail {
			fast = fast.Next
			if fast != tail {
				fast = fast.Next
			}
			slow = slow.Next
		}

		// slow 为中间的指针
		return merge(mergeSort(head, slow), mergeSort(slow, tail))
	}

	return mergeSort(head, nil)
}

func sortList2(head *ListNode) *ListNode {
	arr := make([]*ListNode, 0)
	for head != nil {
		arr = append(arr, head)
		head = head.Next
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})

	for i := 0; i < len(arr)-1; i++ {
		arr[i].Next = arr[i+1]
	}
	arr[len(arr)-1].Next = nil
	return arr[0]
}
