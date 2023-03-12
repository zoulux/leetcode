package main

import . "leetcode"

func main() {

}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}

	return len(generate(1, n))
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
