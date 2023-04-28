package main

import (
	"fmt"
	"github.com/emirpasic/gods/trees/redblacktree"
)

func main() {
	//["DinnerPlates","push","push","push","push","push",
	//"popAtStack","push","push","popAtStack","popAtStack","pop","pop","pop","pop","pop"]
	//[[2],[1],[2],[3],[4],[7],
	//[8],[20],[21],[0],[2],[],[],[],[],[]]

	dp := Constructor(1)
	dp.Push(1)
	dp.Push(2)
	fmt.Println(dp.PopAtStack(1))
	fmt.Println(dp.Pop())
	dp.Push(1)
	dp.Push(2)
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())

}
func main1() {
	//["DinnerPlates","push","push","push","push","push",
	//"popAtStack","push","push","popAtStack","popAtStack","pop","pop","pop","pop","pop"]
	//[[2],[1],[2],[3],[4],[7],
	//[8],[20],[21],[0],[2],[],[],[],[],[]]

	dp := Constructor(2)
	dp.Push(1)
	dp.Push(2)
	dp.Push(3)
	dp.Push(4)
	dp.Push(7)

	fmt.Println(dp.PopAtStack(8))
	dp.Push(20)
	dp.Push(21)
	fmt.Println(dp.PopAtStack(0))
	fmt.Println(dp.PopAtStack(2))
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())
	fmt.Println(dp.Pop())

}

type DinnerPlates struct {
	capacity int     // 栈的容量
	stacks   [][]int // 所有栈
	notFull  *redblacktree.Tree
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{
		capacity: capacity,
		notFull:  redblacktree.NewWithIntComparator(),
	}
}

func (d *DinnerPlates) Push(val int) {
	if s := d.notFull.Size(); s != 0 {
		m := d.notFull.Left()
		idx := m.Key.(int)
		d.stacks[idx] = append(d.stacks[idx], val) // 向不满的里面加一个

		if len(d.stacks[idx]) == d.capacity {
			// 这个时候满了
			d.notFull.Remove(idx)
		}
	} else {
		d.stacks = append(d.stacks, []int{val})
		if d.capacity > 1 {
			d.notFull.Put(len(d.stacks)-1, nil) // 表示这个是不满的
		}
	}

}

func (d *DinnerPlates) Pop() int {
	return d.PopAtStack(len(d.stacks) - 1)

}

func (d *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(d.stacks) {
		return -1
	}

	if len(d.stacks[index]) == 0 {
		return -1
	}

	sh := len(d.stacks[index]) - 1 // 当前栈的高度
	value := d.stacks[index][sh]
	d.stacks[index] = d.stacks[index][:sh] // 高度-1
	d.notFull.Put(index, nil)              // 不管这个栈之前在不在 notFull 中，现在一定是不满的

	for index >= 0 && len(d.stacks[index]) == 0 && len(d.stacks) == index+1 {
		// 高度 -1 之后，高度是否为0，为0则表示可以移除了 ！ 必须在末尾的栈才能被移除，中间的不可以
		d.stacks = d.stacks[:index]

		d.notFull.Remove(index) // 那现在也不在未满栈中了
		index--
	}

	return value
}
