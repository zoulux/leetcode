package main

import (
	"fmt"
	. "leetcode"
	"sort"
)

//107. 二叉树的层序遍历
//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

func main() {
	x := []int{2, 3, 1}
	Reverse(x)
	fmt.Println(x)

	n := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(levelOrderBottom(n))

	n = &TreeNode{
		Val: 3,
	}
	fmt.Println(levelOrderBottom(n))

}

func levelOrderBottom(root *TreeNode) [][]int {

	ans := make([][]int, 0)
	var travel func(root *TreeNode, level int)
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ans) <= level {
			ans = append(ans, []int{})
		}

		ans[level] = append(ans[level], root.Val)
		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}

	travel(root, 0)
	Reverse(ans)
	return ans
}

func Reverse(s interface{}) {
	sort.SliceStable(s, func(i, j int) bool {
		return true
	})
}
