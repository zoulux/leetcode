package leetcode

import "testing"

func climbStairs(n int) int {
	var solve func(n, n1, n2 int) int
	solve = func(n, n1, n2 int) int {

		if n == 1 {
			return n1
		}
		if n == 2 {
			return n2
		}
		n--
		return solve(n, n2, n1+n2)
	}

	return solve(n, 1, 2)
}

func TestClimbStairs(t *testing.T) {
	t.Log(climbStairs2(1))
	t.Log(climbStairs2(2))
	t.Log(climbStairs2(3))
}

func climbStairs2(n int) int {

	p, q := 0, 1
	for i := 1; i <= n; i++ {
		p, q = q, p+q
	}

	return q
}
