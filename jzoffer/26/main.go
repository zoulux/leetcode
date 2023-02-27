package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	fmt.Println(isSubStructure(NewTreeNode(3, 4, 5, 1, 2), NewTreeNode(4, 1)))
	fmt.Println(isSubStructure(NewTreeNode(4, 2, 3, 4, 5, 6, 7, 8, 9), NewTreeNode(4, 8, 9)))
	//[3,4,5,1,2], B = [4,1]
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A == nil {
		return false
	}

	// 比较 B 数和 A 数的节点是否一致
	var isSame func(a, b *TreeNode) bool
	isSame = func(a, b *TreeNode) bool {
		// b 到底了，可以返回 true
		if b == nil {
			return true
		}
		// b 树还没到底，a树已空
		if a == nil && b != nil {
			return false
		}

		// 当前的值相等，再去比较左右子树是否相等
		return a.Val == b.Val && isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
	}

	// 如果两个根结点相等，再去具体比较
	if A.Val == B.Val && isSame(A, B) {
		return true
	}

	// 再去左右两边看看
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}
