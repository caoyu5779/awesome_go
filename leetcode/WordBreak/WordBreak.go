package WordBreak

import (
	"sort"
	"fmt"
)

func WordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0{
		return false
	}

	dict := make(map[string]bool, len(wordDict))
	length := make(map[int]bool, len(wordDict))

	for _,w := range wordDict {
		length[len(w)] = true
		dict[w] = true
	}

	fmt.Println(length)
	fmt.Println(dict)

	sizes := make([]int, 0,len(length))
	for k := range length {
		sizes = append(sizes, k)
	}

	sort.Ints(sizes)
    //[3,5]
	dp := make([]bool, len(s) + 1)

	dp [0] = true
	n := len(s)

	for i := 0; i <= n ; i ++ {
		if !dp[i] {
			continue
		}
		for _, size := range sizes {
			//3 < n
			fmt.Println(i+size)
			if i+size <= n {
				//dp[5] = dp[5] || dict[s[0:5]]
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]
			}
		}
	}

	return dp[n]
}