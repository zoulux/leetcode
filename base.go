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

// 0代表节点为空
func NewTreeNode(values ...int) *TreeNode {
	return createBinaryTree(0, values)
}

func createBinaryTree(i int, nums []int) *TreeNode {
	if nums[i] == 0 {
		return nil
	}
	tree := &TreeNode{Val: nums[i]}
	//左节点的数组下标为1,3,5...2*i+1
	if i < len(nums) && 2*i+1 < len(nums) {
		tree.Left = createBinaryTree(2*i+1, nums)
	}
	//右节点的数组下标为2,4,6...2*i+2
	if i < len(nums) && 2*i+2 < len(nums) {
		tree.Right = createBinaryTree(2*i+2, nums)
	}
	return tree
}
