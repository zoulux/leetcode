package main

import "fmt"

func main() {

	fmt.Println(verifyPostorder([]int{1, 6, 3, 2, 5}))
	fmt.Println(verifyPostorder([]int{1, 3, 2, 6, 5}))
}

func verifyPostorder(postorder []int) bool {

	var verify func(left, right int) bool

	verify = func(left, right int) bool {
		if left >= right {
			return true
		}
		rootValue := postorder[right]
		k := left
		for k < right && postorder[k] < rootValue {
			k++
		}

		// k 为临界值，比 root 大
		// 比 k 小的都是左子树的
		// 大于等于 k 的位置都是右子树的

		for i := k; i < right; i++ {
			// 右子树的值原则上应该比 root 大
			if postorder[i] < rootValue {
				return false
			}
		}

		if !verify(left, k-1) {
			return false
		}

		if !verify(k, right-1) {
			return false
		}
		return true
	}

	return verify(0, len(postorder)-1)
}
