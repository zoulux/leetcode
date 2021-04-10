package leetcode

import (
	"strconv"
	"strings"
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
func tree2str(t *TreeNode) string {

	if t == nil {
		return ""
	}

	if t.Left == nil && t.Right == nil {
		return strconv.FormatInt(int64(t.Val), 10)
	}

	var sb strings.Builder

	sb.WriteString(strconv.FormatInt(int64(t.Val), 10))

	left := tree2str(t.Left)
	right := tree2str(t.Right)

	sb.WriteString("(")
	sb.WriteString(left)
	sb.WriteString(")")

	if right != "" {
		sb.WriteString("(")
		sb.WriteString(right)
		sb.WriteString(")")
	}

	return sb.String()
}

func TestTree2str(t *testing.T) {
	// tree := GenerateTrees(1, 2, 3, 4)
	tree := GenerateTrees(1, 2, 3, -1, 4)
	t.Log(tree2str(tree))
}
