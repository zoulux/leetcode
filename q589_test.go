package leetcode

import (
	"testing"
)

func preorder(root *Node) (ans []int) {

	if root == nil {
		return
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, cur.Val)
		for i := len(cur.Children) - 1; i >= 0; i-- {
			stack = append(stack, cur.Children[i])
		}

	}
	return
}
func TestPreorder(t *testing.T) {
}
