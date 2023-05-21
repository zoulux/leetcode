package main

import (
	"fmt"
	. "leetcode"
	"math"
)

func main() {

	fmt.Println(maxSumBST(&TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
	}))

	fmt.Println(maxSumBST(&TreeNode{
		Val:   -4,
		Left:  &TreeNode{Val: -2},
		Right: &TreeNode{Val: -5},
	}))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxSumBST(root *TreeNode) int {
	min := func(a, b int) int {
		if b < a {
			return b
		}
		return a
	}
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var ans int

	var dfs func(root *TreeNode) (int, int, int)
	dfs = func(root *TreeNode) (int, int, int) {
		if root == nil {
			return math.MinInt, math.MaxInt, 0
		}
		lmin, lmax, lsum := dfs(root.Left)
		rmin, rmax, rsum := dfs(root.Right)
		x := root.Val
		if x <= lmax || x >= rmin {
			return math.MinInt, math.MaxInt, 0
		}
		s := x + lsum + rsum
		ans = max(ans, s)
		return min(lmin, x), max(rmax, x), s
	}
	dfs(root)
	return ans
}

func maxSumBSTErr(root *TreeNode) int {

	var ans int

	var travel func(root *TreeNode, maxValue, minValue int) (int, bool)
	travel = func(root *TreeNode, maxValue, minValue int) (ret int, is bool) {

		defer func() {
			if ret > ans {
				ans = ret
			}

			if is && root != nil && (root.Val > maxValue || root.Val < minValue) {
				is = false
			}
		}()

		if root == nil {
			return 0, true
		}

		if root.Left == nil && root.Right == nil {
			return root.Val, true
		}

		left, lb := travel(root.Left, root.Val, minValue)
		right, rb := travel(root.Right, root.Val, maxValue)

		if root.Left != nil && root.Left.Val < root.Val && root.Right != nil && root.Right.Val > root.Val {
			if lb && rb {
				return root.Val + left + right, true
			}
			return max(0, left, right), false
		}

		if root.Left != nil && root.Left.Val < root.Val && root.Right == nil {
			if lb {
				return root.Val + left, true
			}
			return left, false
		}

		if root.Left == nil && root.Right != nil && root.Right.Val > root.Val {
			if rb {
				return root.Val + right, true
			}
			return right, false
		}

		return max(0, left, right), false
	}

	travel(root, math.MaxInt, math.MinInt)
	return ans
}

func max(arr ...int) int {
	x := arr[0]
	for _, v := range arr[1:] {
		if v > x {
			x = v
		}
	}
	return x
}
