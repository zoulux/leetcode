package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	//1->4->5, 1->3->4, 2->6

	fmt.Println(mergeKLists([]*ListNode{
		nil,
		{Val: 1},
		//{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		//{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		//{Val: 2, Next: &ListNode{Val: 6}},
	}))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	cur := 0
	for i, v := range lists {
		if lists[cur] == nil || (v != nil && v.Val < lists[cur].Val) {
			cur = i
		}
	}

	h := lists[cur]

	if h == nil {
		return nil
	}

	lists[cur] = lists[cur].Next
	if lists[cur] == nil {
		// 将此节点删除
		left := lists[:cur]
		right := lists[cur+1:]
		lists = append([]*ListNode{}, left...)
		lists = append(lists, right...)
	}

	h.Next = mergeKLists(lists)
	return h
}
