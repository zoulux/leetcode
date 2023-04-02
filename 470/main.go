package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())
	fmt.Println(rand10())

}

func rand10() int {
	//1,7
	// 1,7
	// 1,49
	x := rand7() + rand7()
	return x%10 + 1
}

func rand7() int {
	return 1 + rand.Intn(7)
}
