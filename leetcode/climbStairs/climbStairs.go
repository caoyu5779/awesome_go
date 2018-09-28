package climbStairs

func ClimbStairs(n int) int {
	if n < 2 {
		return n
	}

	rec := make([]int, n+1)
	rec[0], rec[1] = 1, 1

	for i := 2; i <= n; i++ {
		//i = 2 rec[2] = 1 +1 = 2
		// i = 3 rec[3] = 2+1 = 3
		rec[i] = rec[i-1] + rec[i-2]

	}

	return rec[n]
}
