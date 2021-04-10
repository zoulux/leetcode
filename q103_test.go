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

type MTreeNode struct {
	Node  *TreeNode // 节点
	Level int       // 层级
}

func zigzagLevelOrder2(root *TreeNode) (ans [][]int) {

	var dfs func(root *TreeNode, level int)

	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(ans) == level {
			ans = append(ans, []int{root.Val})
		} else {
			if level%2 == 0 {
				ans[level] = append(ans[level], root.Val)
			} else {
				ans[level] = append([]int{root.Val}, ans[level]...)
			}
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return
}

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	queue := []*MTreeNode{
		{
			Node:  root,
			Level: 0,
		},
	}

	ans := [][]int{}

	for len(queue) > 0 {
		cur := queue[0]
		curLevel := cur.Level
		curVal := cur.Node.Val
		queue = queue[1:]

		if len(ans) == curLevel {
			ans = append(ans, []int{curVal})
		} else {
			if curLevel%2 == 0 {
				// 偶数后面追加
				ans[curLevel] = append(ans[curLevel], curVal)

			} else {
				// 奇数前面追加
				ans[curLevel] = append([]int{curVal}, ans[curLevel]...)
			}
		}

		if cur.Node.Left != nil {
			queue = append(queue, &MTreeNode{
				Node:  cur.Node.Left,
				Level: curLevel + 1,
			})
		}

		if cur.Node.Right != nil {
			queue = append(queue, &MTreeNode{
				Node:  cur.Node.Right,
				Level: curLevel + 1,
			})
		}
	}

	return ans
}

func TestZigzagLevelOrder(t *testing.T) {

	// tree := GenerateTrees(1, 2, 3, 4, -1, -1, 5) // [[1],[3,2],[4,5]]
	// tree := GenerateTrees(1, 2, 3, 4, 5, 6, 7)
	tree := GenerateTrees(3, 9, 20, -1, -1, 15, 7)
	t.Log(zigzagLevelOrder(tree))
	t.Log(zigzagLevelOrder2(tree))

}
