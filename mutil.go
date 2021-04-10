package leetcode

import (
	"strconv"
	"strings"
)

// Min is get max num from nums
func Min(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}

// Max is get max num from nums
func Max(nums ...int) int {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) String() string {

	arr := []string{}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		arr = append(arr, strconv.Itoa(top.Val))

		if top.Left != nil {
			stack = append(stack, top.Left)
		}

		if top.Right != nil {
			stack = append(stack, top.Right)
		}
	}
	return strings.Join(arr, ",")
}

func (head *ListNode) String() string {
	cur := head

	arr := []string{}
	for cur != nil {
		arr = append(arr, strconv.Itoa(cur.Val))
		cur = cur.Next
	}
	return strings.Join(arr, ",")
}

// GenerateList GenerateList
func GenerateList(arr []int) *ListNode {
	head := &ListNode{
		Val: -1,
	}

	current := head
	for i := 0; i < len(arr); i++ {
		current.Next = &ListNode{
			Val: arr[i],
		}
		current = current.Next
	}

	return head.Next
}

// GenerateLists GenerateLists
func GenerateLists(arr ...int) *ListNode {
	return GenerateList(arr)
}

// GenerateTree GenerateTree
func GenerateTree(arr []int) *TreeNode {
	if len(arr) <= 0 {
		return nil
	}

	nodes := []*TreeNode{}
	for _, val := range arr {
		nodes = append(nodes, &TreeNode{
			Val: val,
		})
	}

	for i := 0; i < len(arr)/2-1; i++ {
		if node := nodes[2*i+1]; node != nil && node.Val != -1 {
			nodes[i].Left = node
		}

		if node := nodes[2*i+2]; node != nil && node.Val != -1 {
			nodes[i].Right = node
		}
	}
	lastIndex := len(arr)/2 - 1
	if node := nodes[lastIndex*2+1]; node.Val != -1 {
		nodes[lastIndex].Left = node
	}

	if len(arr)%2 == 1 {
		if node := nodes[lastIndex*2+2]; node.Val != -1 {
			nodes[lastIndex].Right = nodes[lastIndex*2+2]
		}
	}

	return nodes[0]
}

// GenerateTrees GenerateTrees
func GenerateTrees(arr ...int) *TreeNode {
	return GenerateTree(arr)
	// return treeCreate(0, arr)
}

//
func treeCreate(i int, arr []int) *TreeNode {
	if arr[i] == -1 {
		return nil
	}
	t := &TreeNode{arr[i], nil, nil}
	if i < len(arr) && 2*i+1 < len(arr) {
		t.Left = treeCreate(2*i+1, arr)
	}
	if i < len(arr) && 2*i+2 < len(arr) {
		t.Right = treeCreate(2*i+2, arr)
	}
	return t
}

// Node  Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}
