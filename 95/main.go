package main

import (
	"fmt"
	. "leetcode"
)

//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。

func main() {
	n := createBinaryTree(4)
	fmt.Println(n)
	//n := generateTrees(3)
	//fmt.Println(n)
}

func createBinaryTree(n int) *TreeNode {
	return create(1, n)
}

func create(start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	//这个地方可以选择从start到end的任何一个值作为根节点
	var val = (start + end) / 2
	return &TreeNode{
		Val:   val,
		Left:  create(start, val-1),
		Right: create(val+1, end),
	}
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return generate(1, n)
}

func generate(start int, end int) []*TreeNode {
	var tree []*TreeNode

	if start > end {
		tree = append(tree, nil)
		return tree
	}

	for i := start; i <= end; i += 1 {
		left := generate(start, i-1)
		right := generate(i+1, end)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees2(n int) []*TreeNode {

	var ans []*TreeNode

	var travel func(selectList map[int]bool, root *TreeNode)

	travel = func(selectList map[int]bool, root *TreeNode) {
		if len(selectList) == n {
			ans = append(ans, root)
			return
		}

		for i := 1; i <= n; i++ {
			if selectList[i] {
				continue
			}
			selectList[i] = true
			travel(selectList, root.Left)
			selectList[i] = false
		}
	}

	travel(make(map[int]bool), new(TreeNode))
	return ans
}
