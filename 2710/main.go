package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(removeTrailingZeros("51230100"))
}

func removeTrailingZeros(num string) string {
	return strings.TrimRight(num, "0")
}
