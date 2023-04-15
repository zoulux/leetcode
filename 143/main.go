package main

import (
	"fmt"
	. "leetcode"
)

func main() {

	h := NewListNode(1, 2, 3, 4)
	reorderList(h)
	fmt.Println(h)

}

func reorderList(head *ListNode) {

	arr := make([]*ListNode, 0)

	for head != nil {
		n := head.Next
		//head.Next = nil
		arr = append(arr, head)
		head = n
	}

	left, right := 0, len(arr)-1
	for true {
		arr[left].Next = arr[right]
		left++
		if left == right {
			break
		}
		arr[right].Next = arr[left]
		right--
	}

	arr[left].Next = nil
}

func reorderList2(head *ListNode) {

	arr := make([]*ListNode, 0)

	for head != nil {
		n := head.Next
		//head.Next = nil
		arr = append(arr, head)
		head = n
	}

	for left, right := 0, len(arr)-1; left < right; left, right = left+1, right-1 {

		if left < right-1 {
			arr[left].Next, arr[right].Next, arr[right-1].Next = arr[right], arr[left].Next, arr[right].Next
		} else {
			if arr[right].Next != nil {
				arr[left].Next, arr[right].Next = arr[right], arr[left].Next
			} else {
				arr[left].Next = arr[right]
			}
		}
	}

}

//leetcode submit region end(Prohibit modification and deletion)
