package main

import "fmt"

func main() {

	ss := Constructor()

	fmt.Println(ss.Next(100))
	fmt.Println(ss.Next(80))
	fmt.Println(ss.Next(60))
	fmt.Println(ss.Next(70))
	fmt.Println(ss.Next(60))
	fmt.Println(ss.Next(75))
	fmt.Println(ss.Next(85))
}

type StockSpanner struct {
	st  [][2]int
	cur int
}

func Constructor() StockSpanner {

	return StockSpanner{st: [][2]int{}}
}

func (this *StockSpanner) Next(price int) int {
	for len(this.st) != 0 && this.st[len(this.st)-1][1] <= price {
		this.st = this.st[:len(this.st)-1]
	}

	pre := -1
	if len(this.st) != 0 {
		pre = this.st[len(this.st)-1][0]
	}

	ans := this.cur - pre

	this.st = append(this.st, [2]int{this.cur, price})
	this.cur++
	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
