package main

import (
	"fmt"
)

// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，每个节点除了有一个 next 指针指向下一个节点，
// 还有一个 random 指针指向链表中的任意节点或者 null。
func main() {

	_1 := &Node{Val: 7}
	_2 := &Node{Val: 13}
	_3 := &Node{Val: 11}
	_4 := &Node{Val: 10}
	_5 := &Node{Val: 1}

	_1.Next = _2
	_2.Next = _3
	_2.Random = _1
	_3.Next = _4
	_3.Random = _5
	_4.Next = _5
	_4.Random = _3
	_5.Random = _1

	fmt.Println(copyRandomList2(_1))

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func (n *Node) String() string {
	var values []int
	for cur := n; cur != nil; cur = cur.Next {

		values = append(values, cur.Val)

	}

	return fmt.Sprintf("%v", values)

}

func copyRandomList(head *Node) *Node {
	mm := make(map[*Node]*Node)

	for cur := head; cur != nil; cur = cur.Next {
		mm[cur] = &Node{Val: cur.Val}
	}

	for cur := head; cur != nil; cur = cur.Next {
		mm[cur].Next = mm[cur.Next]
		mm[cur].Random = mm[cur.Random]
	}

	return mm[head]
}

func copyRandomList2(head *Node) *Node {

	//1 2 3 => 1 1` 2 2` 3 3`
	for cur := head; cur != nil; cur = cur.Next.Next {
		cp := &Node{Val: cur.Val}
		cp.Next = cur.Next
		cur.Next = cp
	}

	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}

	//1 1` 2 2` 3 3`
	dummy := new(Node)
	cp := dummy
	for cur := head; cur != nil; cur = cur.Next {
		cp.Next = cur.Next
		cur.Next = cur.Next.Next
		cp = cp.Next
	}

	return dummy.Next
}
