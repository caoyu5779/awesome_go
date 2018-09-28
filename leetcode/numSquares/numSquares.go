package numSquares

import (
	"math"
)

func NumSquares(n int) int {
	perfects := []int{}

	for i := 1; i*i <= n; i++ {
		perfects = append(perfects, i*i)
	}

	dp := make([]int, n+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}

	for _, p := range perfects {
		for i := p; i < len(dp); i++ {
			if dp[i] > dp[i-p]+1 {
				dp[i] = dp[i-p] + 1
			}
		}
	}

	return dp[n]
}
