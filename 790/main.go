package main

func main() {

}

func numTilings(n int) int {

	mod := int(1e9 + 7)

	switch n {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	}

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = numTilings(1)
	dp[2] = numTilings(2)

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[0]*2) % mod
		dp[0] += dp[i-2] % mod
	}
	return dp[n]
}
