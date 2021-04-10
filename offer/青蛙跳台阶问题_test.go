package offer

import (
	"testing"
)

func numWays(n int) int {
	cache := map[int]int{}

	var ways func(n int) int
	ways = func(n int) int {
		if n == 0 || n == 1 {
			return 1
		} else if n == 2 {
			return 2
		}

		if _, ok := cache[n]; !ok {
			cache[n] = (ways(n-1) + ways(n-2)) % 1000000007
		}
		return cache[n]
	}

	return ways(n)

}

func TestNumWays(t *testing.T) {
	t.Log(numWays(7))
	// t.Log(numWays(100))

}
