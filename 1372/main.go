package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	t := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   5,
					Right: &TreeNode{Val: 7},
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(longestZigZag(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var ans int
	var travel func(root *TreeNode, l, r int)
	travel = func(root *TreeNode, l, r int) {
		if root == nil {
			return
		}

		ans = max(ans, max(l, r))

		if root.Left != nil {
			travel(root.Left, r+1, 0)
		}

		if root.Right != nil {
			travel(root.Right, 0, l+1)
		}
	}

	travel(root, 0, 0)
	return ans
}

func longestZigZag2(root *TreeNode) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	type item struct {
		t    *TreeNode
		left bool
	}
	cache := map[item]int{}

	var travel func(root *TreeNode, left bool) int
	travel = func(root *TreeNode, left bool) int {
		if root == nil {
			return 0
		}

		if v, ok := cache[item{
			t:    root,
			left: left,
		}]; ok {
			return v
		}

		if left {
			if root.Left != nil {
				t := travel(root.Left, false) + 1

				cache[item{
					t:    root,
					left: left,
				}] = t

				return t
			}
			return 0
		}

		if root.Right != nil {
			t := travel(root.Right, true) + 1
			cache[item{
				t:    root,
				left: left,
			}] = t
			return t
		}
		return 0
	}

	var ans int
	var travel2 func(root *TreeNode)
	travel2 = func(root *TreeNode) {
		if root == nil {
			return
		}

		ans = max(ans, max(
			travel(root, true),
			travel(root, false),
		))

		travel2(root.Left)
		travel2(root.Right)
	}

	travel2(root)

	return ans
}
