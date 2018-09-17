package minPathSum

import "fmt"

func MinPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	dp := make([][]int , m)

	for i := range dp {
		dp[i] = make([]int, n)
	}
	//这样就做好了一个空的矩阵
	dp[0][0] = grid[0][0]

	for i := 1; i < m ; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
		//dp1,0 = grid 1,0+dp0,0 = 1+1 = 2
		//dp2,0 = grid 2,0+dp1,0 = 4+1 = 6
	}

	fmt.Println(dp)

	for j := 1; j < n ; j ++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}

	for i := 1; i < m ; i ++ {
		for j := 1; j < n ; j ++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}

	return dp[m-1][n-1]
}

func min(i,j int) int {
	if i > j {
		return j
	}

	return i
}
