package leetcode

import (
	"testing"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// type BSTIterator struct {
// 	arr []int
// }

// func iterator(root *TreeNode) (arr []int) {
// 	if root == nil {
// 		return
// 	}

// 	left := iterator(root.Left)
// 	arr = append(arr, root.Val)
// 	right := iterator(root.Right)

// 	return append(append(left, arr...), right...)
// }

// func Constructor22(root *TreeNode) BSTIterator {
// 	return BSTIterator{arr: iterator(root)}
// }

// func (this *BSTIterator) Next() int {
// 	v := this.arr[0]
// 	this.arr = this.arr[1:]
// 	return v
// }

// func (this *BSTIterator) HasNext() bool {
// 	return len(this.arr) > 0
// }

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor22(root *TreeNode) BSTIterator {

	stack := []*TreeNode{}
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	return BSTIterator{stack}
}

func (this *BSTIterator) Next() int {
	top := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	right := top.Right
	for right != nil {
		this.stack = append(this.stack, right)
		right = right.Left
	}

	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func TestBSTIterator(t *testing.T) {

	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	bst := Constructor22(root)
	t.Log(bst.Next())
	t.Log(bst.Next())
	t.Log(bst.HasNext())
	t.Log(bst.Next())
	t.Log(bst.HasNext())
	t.Log(bst.Next())
	t.Log(bst.HasNext())
	t.Log(bst.Next())
	t.Log(bst.HasNext())

}

// class BSTIterator(object):

//     def __init__(self, root):
//         self.stack = []
//         while root:
//             self.stack.append(root)
//             root = root.left

//     def next(self):
//         cur = self.stack.pop()
//         node = cur.right
//         while node:
//             self.stack.append(node)
//             node = node.left
//         return cur.val

//     def hasNext(self):
//         return len(self.stack) > 0
