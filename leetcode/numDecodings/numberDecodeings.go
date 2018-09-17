package numDecodings

func NumDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, one(s[0:1])

	for i:=2; i<=n;i++ {
		w1,w2 := one(s[i-1:i]), two(s[i-2:i])

		//dp[2] = dp[1]*1+dp[0]*1 = 2
		dp[i] = dp[i-1] * w1 + dp[i-2] * w2

		if dp[i] == 0{
			return 0
		}
	}

	return dp[n]
}

/*
	检查s是否为 合格的单字符
	len(s) == 1
	合格 返回1
	不合格 返回0
*/
func one(s string) int {
	if s == "0" {
		return 0
	}

	return 1
}

/*
	检查s是否为合格的双字符
	len(s) == 2
	合格返回1
	不合格 返回0
	NOTICE : GO 强语言类型
 */
func two(s string) int {
	switch s[0] {
	case '1':
		return 1
	case '2':
		if s[1] == '7' || s[1] == '8' || s[1] == '9' {
			return 0
		}
		return 1
	default:
		return 0
	}
}