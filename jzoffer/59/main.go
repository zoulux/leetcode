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
	values    []int
	maxValues []int
}

func Constructor() MaxQueue {
	return MaxQueue{values: []int{}}
}

func (this *MaxQueue) Max_value() int {
	return this.maxValues[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.values = append(this.values, value)
	if len(this.maxValues) == 0 {
		this.maxValues = append(this.maxValues, value)
	} else {
		_max := this.maxValues[len(this.maxValues)-1]
		if value > _max {
			this.maxValues = append(this.maxValues, value)
		} else {
			this.maxValues = append(this.maxValues, _max)
		}
	}
}

func (this *MaxQueue) Pop_front() int {

	if len(this.values) == 0 {
		return -1
	}
	x := this.values[0]

	this.values = this.values[1:]
	return x
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */
