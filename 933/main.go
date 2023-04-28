package main

import "fmt"

func main() {

	//[null,1,2,1,2,3]

	rc := Constructor()
	fmt.Println(rc.Ping(642))
	fmt.Println(rc.Ping(1849))
	fmt.Println(rc.Ping(4921))
	fmt.Println(rc.Ping(5936))
	fmt.Println(rc.Ping(5957))

}

type RecentCounter struct {
	s []int
}

func Constructor() RecentCounter {

	return RecentCounter{
		s: []int{},
	}

}

func (this *RecentCounter) Ping(t int) int {

	for len(this.s) > 0 && this.s[0] < t-3000 {
		this.s = this.s[1:]
	}

	this.s = append(this.s, t)

	return len(this.s)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
