package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	res := pathSum(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{
				Val:   2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val:   -3,
			Right: &TreeNode{Val: 11},
		},
	}, 8)

	//res := pathSum(&TreeNode{
	//	Val: 1,
	//}, 1)
	fmt.Println(res)
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	mm := make(map[int]int)
	mm[0] = 1

	var ans int
	var travel func(root *TreeNode, val int)
	travel = func(root *TreeNode, val int) {
		ans += mm[val-targetSum]
		mm[val]++

		if root.Left != nil {
			travel(root.Left, val+root.Left.Val)
		}

		if root.Right != nil {
			travel(root.Right, val+root.Right.Val)
		}

		mm[val]--
	}

	travel(root, root.Val)
	return ans
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum2(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	mm := make(map[int]int)
	mm[0] = 1

	var ans int
	var travel func(root *TreeNode, val int)
	travel = func(root *TreeNode, val int) {
		ans += mm[val-targetSum]
		mm[val]++

		if root.Left != nil {
			travel(root.Left, val+root.Left.Val)
		}

		if root.Right != nil {
			travel(root.Right, val+root.Right.Val)
		}

		mm[val]--
	}

	travel(root, root.Val)
	return ans
}
