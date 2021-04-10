package leetcode

type MyQueue struct {
	list []int
}

// /** Initialize your data structure here. */
// func Constructor() MyQueue {
// 	return MyQueue{list: []int{}}
// }

// /** Push element x to the back of queue. */
// func (this *MyQueue) Push(x int) {
// 	this.list = append(this.list, x)
// }

// /** Removes the element from in front of queue and returns that element. */
// func (this *MyQueue) Pop() int {
// 	ans := this.list[0]
// 	this.list = this.list[1:]

// 	return ans
// }

// /** Get the front element. */
// func (this *MyQueue) Peek() int {
// 	ans := this.list[0]
// 	return ans
// }

// /** Returns whether the queue is empty. */
// func (this *MyQueue) Empty() bool {
// 	return len(this.list) == 0
// }
