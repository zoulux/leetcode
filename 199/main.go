package main

import . "leetcode"

func main() {

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
func rightSideView(root *TreeNode) []int {
	var ans [][]int
	var travel func(root *TreeNode, level int)
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if len(ans) < level {
			ans = append(ans, []int{})
		}
		ans[level] = append(ans[level], root.Val)
		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}
	travel(root, 0)

	var res []int
	for _, v := range ans {
		res = append(res, v[len(v)-1])
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
