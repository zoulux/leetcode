package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	ch := make(chan int)
	go func() {
		for {
			select {
			case v, ok := <-ch:
				fmt.Println(v, "1")
				if !ok {
					return
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case v, ok := <-ch:
				fmt.Println(v, "2")
				if !ok {
					return
				}
			}
		}
	}()

	//go func() {
	//	time.Sleep(time.Millisecond)
	//	close(ch)
	//}()
	//
	ch <- 100

	close(ch)

	time.Sleep(time.Millisecond * 2)

	//fmt.Println(sumSubarrayMins([]int{
	//	3, 1, 2, 4,
	//}))
}

func sumSubarrayMins(arr []int) int {
	sort.Ints(arr)

	s := 0
	sum := 0
	// [1,2,3,4]
	for _, v := range arr {
		sum += v
		s = s + sum

	}
	return s
	//	 1
	// [1] => 1

	// 12
	// [1] => 1
	// [2] => 2
	// [12] => 1
	// 1*2+2

	// 123
	// [1] => 1
	// [2] => 2
	// [12] => 1
	// [3] => 3
	// [13] => 1
	// [23] => 2
	// 1*3+2*2+3*1
	//1*2+2+ 1+2+3

	// 1234
	// 1   1
	// 2   2
	// 3   3
	// 4   4
	// 12  1
	// 13  1
	// 14  1
	// 23  2
	// 24  2
	// 34  3
	// 123 1
	// 234 2
	// 1234 1
	//  1*6 + 2*4 + 3*2 +  4*1
}
