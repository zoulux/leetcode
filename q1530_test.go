package leetcode

import (
	"fmt"
	. "mutil"
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
func countPairs(root *TreeNode, distance int) int {

	ans := 0

	// filterFun := func(k *TreeNode, v int) bool {
	// 	return v < distance
	// }

	var dfs func(root *TreeNode) map[*TreeNode]int

	dfs = func(root *TreeNode) map[*TreeNode]int {
		if root == nil {
			return map[*TreeNode]int{}
		}

		// 叶子节点
		if root.Left == nil && root.Right == nil {
			return map[*TreeNode]int{root: 0}
		}

		leftLeaf := dfs(root.Left)
		rightLeaf := dfs(root.Right)

		for k, v := range rightLeaf {
			if v >= distance {
				delete(rightLeaf, k)
			} else {
				rightLeaf[k] = v + 1
			}
		}

		for lk, lv := range leftLeaf {

			if lv >= distance {
				delete(leftLeaf, lk)
				continue
			} else {
				leftLeaf[lk] = lv + 1
				lv++
			}

			for _, rv := range rightLeaf {
				if lv+rv <= distance {
					ans++
				}
			}
		}

		for k, v := range rightLeaf {
			leftLeaf[k] = v
		}

		return leftLeaf
	}
	dfs(root)
	return ans
}

func filter(left map[*TreeNode]int, f func(k *TreeNode, v int) bool) {
	for k, v := range left {
		if !f(k, v) {
			delete(left, k)
		}
	}
}

func add(left map[*TreeNode]int) {
	for k, v := range left {
		left[k] = v + 1
	}
}

func TestCountPairs(t *testing.T) {
	// tree := GenerateTrees(1,2,3,-1,4)
	tree := GenerateTrees(7, 1, 4, 6, -1, 5, 3, -1, -1, -1, -1, -1, 2)
	t.Log(countPairs(tree, 3))
	t.Log(countPairs3(tree, 3))
}

func countPairs2(root *TreeNode, distance int) (ans int) {
	var dfs func(*TreeNode) []int
	dfs = func(root *TreeNode) (n []int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			return []int{1}
		}
		ln := dfs(root.Left)
		rn := dfs(root.Right)
		for i := range ln {
			n = append(n, ln[i]+1)
			for j := range rn {
				if ln[i]+rn[j] <= distance {
					ans++
				}
			}
		}

		for j := range rn {
			n = append(n, rn[j]+1)
		}

		return
	}
	dfs(root)
	return
}

func traverse1(root *TreeNode) {
	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		fmt.Print(root.Val)

		dfs(root.Left)
		dfs(root.Right)

	}
	fmt.Println("前序遍历")
	dfs(root)
	fmt.Println()

}

func traverse2(root *TreeNode) {
	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		fmt.Print(root.Val)

		dfs(root.Right)

	}
	fmt.Println("中序遍历")

	dfs(root)
	fmt.Println()

}

func traverse3(root *TreeNode) {
	var dfs func(*TreeNode)

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)

		fmt.Print(root.Val)
	}
	fmt.Println("后序遍历")
	dfs(root)
	fmt.Println()
}

func TestTraverse(t *testing.T) {
	tree := GenerateTrees(1, 2, 3, -1, 4)
	traverse1(tree)
	traverse2(tree)
	traverse3(tree)
}

func countPairs3(root *TreeNode, distance int) int {

	ans := 0

	var dfs func(root *TreeNode) []int

	dfs = func(root *TreeNode) (n []int) {
		if root == nil {
			return n
		}

		// 叶子节点
		if root.Left == nil && root.Right == nil {
			n = append(n, 0)
			return n
		}

		leftLeaf := dfs(root.Left)
		rightLeaf := dfs(root.Right)

		for k, v := range leftLeaf {
			leftLeaf[k] = v + 1
		}

		for k, v := range rightLeaf {
			rightLeaf[k] = v + 1
		}

		for _, lv := range leftLeaf {
			for _, rv := range rightLeaf {
				if lv+rv <= distance {
					ans++
				}
			}
		}

		return append(leftLeaf, rightLeaf...)
	}
	dfs(root)
	return ans
}

func TestArrayDel(t *testing.T) {
	leftLeaf := []int{1, 2, 3, 4}
	distance := 2
	for i := 0; i < len(leftLeaf); i++ {
		v := leftLeaf[i]
		if v <= distance {
			leftLeaf[i] = v
		} else {
			leftLeaf = append(leftLeaf[:i], leftLeaf[i+1:]...)
			i--
		}
	}

	t.Log(leftLeaf)
}
