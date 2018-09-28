package longestPalindrome

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	begin, maxLen := 0, 1

	for i:=0; i < len(s); {
		if len(s) - i < maxLen/2{
			break
		}

		b,e := i, i

		for e < len(s) - 1 && s[e+1] == s[e] {
			e++
		}

		//e=9
		i = e +1

		//i=10
		//for e < 10 && b > 0 && s[10] == s[0]
		for e < len(s) - 1 && b > 0 && s[e+1] == s[b-1] {
			e++
			b--
		}
		//11
		newLen := e+1-b

		if newLen > maxLen {
			//begin=0
			begin = b
			//11
			maxLen = newLen
		}
	}

	return s[begin:begin+maxLen]
}
