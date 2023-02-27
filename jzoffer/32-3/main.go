package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(levelOrder(&TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 8}},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}))
}

func levelOrder1(root *TreeNode) [][]int {
	var arr [][]int
	var travel func(root *TreeNode, level int)
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(arr) <= level {
			arr = append(arr, []int{})
		}
		if level%2 == 0 {
			arr[level] = append(arr[level], root.Val)

		} else {
			arr[level] = append([]int{root.Val}, arr[level]...)

		}

		fmt.Println("LEVEL", level, arr)
		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}

	travel(root, 0)

	return arr
}

func levelOrder(root *TreeNode) [][]int {
	var arr [][]int
	var travel func(root *TreeNode, level int)
	travel = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(arr) <= level {
			arr = append(arr, []int{})
		}
		arr[level] = append(arr[level], root.Val)

		fmt.Println("LEVEL", level, arr)
		travel(root.Left, level+1)
		travel(root.Right, level+1)
	}

	travel(root, 0)

	reverse := func(arr []int) {
		left, right := 0, len(arr)-1
		for left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	for i := range arr {
		if i%2 == 1 {
			reverse(arr[i])
		}
	}

	return arr
}
