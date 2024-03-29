package main

import (
	"fmt"
	. "leetcode"
	"math"
)

func main() {
	//NewTreeNode(-10, 9, 20, nil, nil, 15, 7)

	n1 := maxPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	})

	fmt.Println(n1)

	n1 = maxPathSum(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	})
	fmt.Println(n1)

	n1 = maxPathSum(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: -1,
		},
	})
	fmt.Println(n1)

	n1 = maxPathSum(&TreeNode{
		Val: -2,
		Left: &TreeNode{
			Val: -1,
		},
	})
	fmt.Println(n1)
}

func maxPathSum(root *TreeNode) int {

	max := func(arr ...int) int {
		m := arr[0]
		for _, v := range arr[1:] {
			if v > m {
				m = v
			}
		}
		return m
	}

	ans := math.MinInt

	var _maxPathSum func(root *TreeNode) int
	_maxPathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		t1 := _maxPathSum(root.Left)
		t2 := _maxPathSum(root.Right)

		t := max(root.Val, root.Val+t1, root.Val+t2, root.Val+t1+t2)
		// 自已作为中间节点计算 （自己，自己+左边，自已+右边，自己+左边+右边）
		ans = max(ans, t)
		// 自己作为节点贡献 (自由，自己+左边，自己+右边)
		return max(root.Val, root.Val+t1, root.Val+t2)
	}
	_maxPathSum(root)
	return ans
}

func maxPathSum2(root *TreeNode) int {
	ans := math.MinInt
	var _maxPathSum func(root *TreeNode) int
	_maxPathSum = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := _maxPathSum(root.Left)
		r := _maxPathSum(root.Right)
		// 当前节点的最大路径： max(自己，自己+左边，自己+右边，自己 + 左边 + 右边）
		ans = max(ans, max(0, l, r, l+r)+root.Val)
		//当前节点作为子节点时的贡献：max(自己，自己+左边，自己+右边）
		return max(0, l, r) + root.Val
	}

	_maxPathSum(root)
	return ans
}

func max(values ...int) int {
	ans := math.MinInt
	for _, v := range values {
		if v > ans {
			ans = v
		}
	}
	return ans
}
