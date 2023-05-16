package main

import (
	"strconv"
	"strings"
)

func main() {

}

func queryString(s string, n int) bool {
	for i := 1; i <= n; i++ {
		if !strings.Contains(s, strconv.FormatInt(int64(i), 2)) {
			return false
		}
	}
	return true
}
