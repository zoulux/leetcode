package main

import . "leetcode"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	var _getIntersectionNode func(headA, headB *ListNode) *ListNode

	_getIntersectionNode = func(a, b *ListNode) *ListNode {
		if a == b {
			return a
		}

		if a == nil {
			return _getIntersectionNode(headB, b)
		}

		if b == nil {
			return _getIntersectionNode(a, headA)
		}
		return nil
	}

	return _getIntersectionNode(headA, headB)
}

//leetcode submit region end(Prohibit modification and deletion)
