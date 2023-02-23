package main

import "fmt"

type Stack struct {
	values []int
}

func NewStack() *Stack {
	return &Stack{
		values: []int{},
	}
}

func (s *Stack) Push(v int) {
	s.values = append(s.values, v)
}
func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) Pop() int {
	if len(s.values) > 0 {
		h := s.values[s.Len()-1]
		s.values = s.values[:s.Len()-1]
		return h
	}
	return -1
}

type CQueue struct {
	s1 *Stack
	s2 *Stack
}

func Constructor() CQueue {
	return CQueue{
		s1: NewStack(),
		s2: NewStack(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.s1.Push(value)
}

func (this *CQueue) DeleteHead() int {

	if this.s2.Len() == 0 {
		for {
			k := this.s1.Pop()
			if k == -1 {
				break
			}
			this.s2.Push(k)
		}
	}
	return this.s2.Pop()
}

func main2() {
	q := Constructor()
	q.AppendTail(8)
	q.AppendTail(20)
	q.AppendTail(1)
	q.AppendTail(11)
	q.AppendTail(2)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}

func main() {

	q := Constructor()

	fmt.Println(q.DeleteHead())
	q.AppendTail(12)
	fmt.Println(q.DeleteHead())

	q.AppendTail(10)
	q.AppendTail(9)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	q.AppendTail(20)
	fmt.Println(q.DeleteHead())
	q.AppendTail(1)
	q.AppendTail(8)
	q.AppendTail(20)
	q.AppendTail(1)
	q.AppendTail(11)
	q.AppendTail(2)
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())

}
