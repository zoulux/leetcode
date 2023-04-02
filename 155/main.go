package main

import "fmt"

func main() {

	fmt.Println()

}

// leetcode submit region begin(Prohibit modification and deletion)
type MinStack struct {
	values    []int
	minValues []int
}

func Constructor() MinStack {
	return MinStack{
		values:    []int{},
		minValues: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.values = append(this.values, val)
	if len(this.minValues) > 0 {
		if m := this.minValues[len(this.minValues)-1]; m < val {
			this.minValues = append(this.minValues, m)
			return
		}
	}
	this.minValues = append(this.minValues, val)
}

func (this *MinStack) Pop() {

	n := len(this.values)
	if len(this.values) > 0 {
		this.values = this.values[:n-1]
		this.minValues = this.minValues[:n-1]
	}

}

func (this *MinStack) Top() int {
	return this.values[len(this.values)-1]
}

func (this *MinStack) GetMin() int {
	return this.minValues[len(this.minValues)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
//leetcode submit region end(Prohibit modification and deletion)
