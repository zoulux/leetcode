package leetcode

import (
	"container/heap"
	. "mutil"
	"sort"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {

	head := new(ListNode)

	var merge func(lists []*ListNode, last *ListNode)

	merge = func(lists []*ListNode, last *ListNode) {

		var min *ListNode
		minIndex := -1

		for index, list := range lists {
			if list != nil {
				if min == nil || list.Val < min.Val {
					min = list
					minIndex = index
				}
			}
		}

		if minIndex != -1 {
			last.Next = min
			last = last.Next

			if min.Next != nil {
				lists[minIndex] = min.Next
			} else {
				lists = append(lists[0:minIndex], lists[minIndex+1:]...)
			}

			merge(lists, last)
		}
	}

	merge(lists, head)
	return head.Next
}

func TestMergeKLists(t *testing.T) {

	t.Log(mergeKLists([]*ListNode{
		GenerateLists(1, 4, 5),
		GenerateLists(),
		GenerateLists(1, 3, 4),
		GenerateLists(2, 6),
	},
	))
}

func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		// [[0],[1],[2],[0,1]]
		lists = append(lists, merge(lists[0], lists[1]))
		lists = lists[2:]
	}
	return lists[0]
}

func merge(one *ListNode, two *ListNode) *ListNode {

	root := new(ListNode)
	cur := root
	for one != nil && two != nil {
		if one.Val < two.Val {
			cur.Next = one
			one = one.Next
		} else {
			cur.Next = two
			two = two.Next
		}
		cur = cur.Next
	}
	if one != nil {
		cur.Next = one
	} else {
		cur.Next = two
	}
	return root.Next
}

func mergeKLists3(lists []*ListNode) *ListNode {

	arr := []*ListNode{}
	for _, list := range lists {
		for list != nil {
			arr = append(arr, &ListNode{
				Val: list.Val,
			})
			list = list.Next
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Val < arr[j].Val
	})

	dummyHead := new(ListNode)
	tail := dummyHead
	for _, item := range arr {
		tail.Next = item
		tail = tail.Next

	}
	return dummyHead.Next
}

func TestMergeKLists3(t *testing.T) {

	// [[],[],[],[-9,-6,-5,-2,4,4],[-9,-6,-5,-2,-1,3],[-2,-1,0],[-6],[-8,-1,0,2]]

	t.Log(mergeKLists4([]*ListNode{
		GenerateLists(-6, -4, 0, 0, 4),
		GenerateLists(),
		GenerateLists(-4, -2, -1, 1, 2, 3),
	},
	))
}

func mergeKLists4(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := nodeHeap{}
	for _, n := range lists {
		if n != nil {
			heap.Push(&h, n)
		}
	}
	dummy := new(ListNode)
	p := dummy
	for len(h) > 0 {
		cur := heap.Pop(&h).(*ListNode)
		if cur.Next != nil {
			heap.Push(&h, cur.Next)
		}
		p.Next = cur
		p = p.Next
	}
	return dummy.Next
}

type nodeHeap []*ListNode

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
