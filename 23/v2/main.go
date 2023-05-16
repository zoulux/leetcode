package main

import (
	"fmt"
	. "leetcode"
	"sort"
)

func main() {

	lists := []int{1, 1, 0, 2, 3}

	left, right := 0, 0
	for ; right < len(lists); right++ {
		if lists[right] != 0 {
			lists[right], lists[left] = lists[left], lists[right]
			left++
		}
	}
	fmt.Println(lists[:left])

	fmt.Println(mergeKLists([]*ListNode{
		NewListNode(),
		NewListNode(1),
	}))
	fmt.Println(mergeKLists([]*ListNode{
		NewListNode(1, 4, 5),
		NewListNode(1, 3, 4),
		NewListNode(2, 6),
	}))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	left, right := 0, 0
	for ; right < len(lists); right++ {
		if lists[right] != nil {
			lists[right], lists[left] = lists[left], lists[right]
			left++
		}
	}
	lists = lists[:left]
	if len(lists) == 0 {
		return nil
	}

	sort.Slice(lists, func(i, j int) bool {
		return lists[i].Val < lists[j].Val
	})

	head := lists[0]

	lists[0] = lists[0].Next

	if lists[0] == nil {
		lists = lists[1:]
	}

	head.Next = mergeKLists(lists)
	return head
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	cur := -1
	for i, l := range lists {
		if cur == -1 && l != nil {
			cur = i
		}

		if cur != -1 && l != nil && l.Val < lists[cur].Val {
			cur = i
		}
	}

	if cur == -1 {
		return nil
	}

	head := lists[cur]
	if lists[cur] != nil {
		lists[cur] = lists[cur].Next
	}

	if lists[cur] == nil {
		lists = append(lists[:cur], lists[cur+1:]...)
	}

	head.Next = mergeKLists(lists)
	return head
}
