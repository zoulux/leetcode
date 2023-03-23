package main

func main() {

	//["MaxQueue","push_back","push_back","max_value","pop_front","max_value"]
	//	[[],[1],[2],[],[],[]]
	//	输出: [null,null,null,2,1,2]

	que := Constructor()
	que.Push_back(1)
	que.Max_value()
	que.Push_back(2)
	que.Max_value()
	que.Pop_front()
	que.Max_value()

}

type MaxQueue struct {
	values []int
	dq     []int
}

func Constructor() MaxQueue {
	return MaxQueue{values: []int{}, dq: []int{}}
}

func (this *MaxQueue) Max_value() int {
	if len(this.dq) == 0 {
		return -1
	}
	return this.dq[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.values = append(this.values, value)

	for len(this.dq) != 0 && this.dq[len(this.dq)-1] < value {
		this.dq = this.dq[:len(this.dq)-1]
	}
	this.dq = append(this.dq, value)

}

func (this *MaxQueue) Pop_front() int {
	if len(this.values) == 0 {
		return -1
	}
	x := this.values[0]

	this.values = this.values[1:]

	if len(this.dq) != 0 && x == this.dq[0] {
		this.dq = this.dq[1:]
	}

	// 3 5 4
	// 3 => 5 => 5 4

	return x
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
