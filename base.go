package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var arr []int
	h := l
	for h != nil {
		arr = append(arr, h.Val)
		h = h.Next
	}

	return fmt.Sprint(arr)
}

func NewListNode(values ...int) *ListNode {
	dummy := new(ListNode)

	head := dummy
	for _, v := range values {
		head.Next = &ListNode{
			Val: v,
		}
		head = head.Next
	}
	return dummy.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func (t *TreeNode) String() string {
//
//}

func NewTreeNode(values ...int) *TreeNode {
	return &TreeNode{}
}
