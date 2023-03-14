package main

import "fmt"

func main() {

	fmt.Println(hammingWeight(11))

}
func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
		if num == 0 {
			break
		}
	}
	return count
}
