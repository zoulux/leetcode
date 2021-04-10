package leetcode

import (
	"testing"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	var head, last *ListNode

	for l1 != nil || l2 != nil {

		if l1 != nil && l2 != nil {

			if last == nil {

				if l1.Val < l2.Val {
					last = &ListNode{
						Val: l1.Val,
					}
					l1 = l1.Next
				} else {
					last = &ListNode{
						Val: l2.Val,
					}
					l2 = l2.Next
				}

			} else {

				if l1.Val < last.Val {
					last.Next = &ListNode{
						Val: l1.Val,
					}
					l1 = l1.Next
				} else if l2.Val < last.Val {
					last.Next = &ListNode{
						Val: l2.Val,
					}
					l2 = l2.Next
				} else {
					if l1.Val < l2.Val {
						last.Next = &ListNode{
							Val: l1.Val,
						}
						l1 = l1.Next
					} else {
						last.Next = &ListNode{
							Val: l2.Val,
						}
						l2 = l2.Next
					}
				}

			}

		} else if l1 != nil {
			if last == nil {
				last = &ListNode{
					Val: l1.Val,
				}
			} else {
				last.Next = &ListNode{
					Val: l1.Val,
				}
			}
			l1 = l1.Next

		} else if l2 != nil {

			if last == nil {
				last = &ListNode{
					Val: l2.Val,
				}
			} else {
				last.Next = &ListNode{
					Val: l2.Val,
				}
			}
			l2 = l2.Next
		}

		if head == nil {
			head = last
		}

		if last.Next != nil {
			last = last.Next
		}

	}

	return head
}

func TestMergeTwoLists(t *testing.T) {
	l1 := GenerateList([]int{1, 2, 4})
	l2 := GenerateList([]int{1, 3, 4})

	t.Log(mergeTwoLists(l1, l2))

}
