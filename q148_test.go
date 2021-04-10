package leetcode

import (
	"fmt"
	. "mutil"
	"testing"
	"time"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	var sort func(head, tail *ListNode) *ListNode

	merge := func(head1, head2 *ListNode) *ListNode {

		dummy := new(ListNode)
		tmp, tmp1, tmp2 := dummy, head1, head2

		for tmp1 != nil && tmp2 != nil {
			if tmp1.Val <= tmp2.Val {
				tmp.Next = tmp1
				tmp1 = tmp1.Next
			} else {
				tmp.Next = tmp2
				tmp2 = tmp2.Next
			}
			tmp = tmp.Next
		}

		if tmp1 != nil {
			tmp.Next = tmp1
		} else if tmp2 != nil {
			tmp.Next = tmp2
		}

		return dummy.Next
	}

	sort = func(head, tail *ListNode) *ListNode {
		if head == nil {
			return head
		}

		if head.Next == tail {
			head.Next = nil
			return head
		}

		fast, slow := head, head
		for fast != tail {
			fast = fast.Next
			slow = slow.Next
			if fast != tail {
				fast = fast.Next
			}
		}

		return merge(
			sort(head, slow),
			sort(slow, tail),
		)
	}

	return sort(head, nil)
}

func TestSortList(t *testing.T) {
	// li := GenerateLists(4, 1, 2, 3)
	// t.Log(sortList(li))

}

func Test1(t *testing.T) {
	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	add, sub := adder(), adder()

	for i := 0; i < 4; i++ {
		fmt.Println(add(i), sub(-i))
	}
}

func Test2(t *testing.T) {
	sum := 0

	adder := func() func(int) int {
		return func(x int) int {
			sum += x
			return sum
		}
	}

	add, sub := adder(), adder()

	for i := 0; i < 4; i++ {
		fmt.Println(add(i), sub(-i))
	}
}

func Test3(t *testing.T) {
	a := 0
	go func() {
		for {
			a++
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(a)
}

func Test4(t *testing.T) {
	ch := make(chan int, 0)

	a := 0
	go func() {
		for {
			select {
			case <-ch:
				fmt.Println("ch")
			default:
				a++
			}
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(a)
}
