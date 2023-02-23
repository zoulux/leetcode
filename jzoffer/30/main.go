package main

import "fmt"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.Min()) //  --> 返回 -3.
	minStack.Pop()
	fmt.Println(minStack.Top()) //    --> 返回 0.
	fmt.Println(minStack.Min()) // --> 返回 -2.
}

type MinStack struct {
	values    []int
	minValues []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{values: []int{}}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	if len(this.minValues) == 0 {
		this.minValues = append(this.minValues, x)
	} else {
		miniTop := this.minValues[len(this.minValues)-1]
		if x < miniTop {
			this.minValues = append(this.minValues, x)
		} else {
			this.minValues = append(this.minValues, miniTop)
		}
	}
}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
	this.minValues = this.minValues[:len(this.minValues)-1]
}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) Min() int {
	return this.minValues[len(this.minValues)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
