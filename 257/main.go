package main

import (
	"fmt"
	. "leetcode"
	"strconv"
	"strings"
)

func main() {

	t := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	fmt.Println(binaryTreePaths(&t))

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {

	var ans [][]int
	var travel func(root *TreeNode, s []int)
	travel = func(root *TreeNode, s []int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			s = append(s, root.Val)
			ans = append(ans, append([]int{}, s...))
			return
		}
		travel(root.Left, append(s, root.Val))
		travel(root.Right, append(s, root.Val))
	}

	travel(root, []int{})

	var res []string

	for _, v := range ans {

		var tmp []string
		for _, vv := range v {
			tmp = append(tmp, strconv.Itoa(vv))
		}
		res = append(res, strings.Join(tmp, "->"))
	}

	return res
	//return ans
}

//leetcode submit region end(Prohibit modification and deletion)
