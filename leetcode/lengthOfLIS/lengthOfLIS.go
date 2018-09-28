package lengthOfLIS

import (
	"sort"
)

func LengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))

	for _, n := range nums {
		at := sort.SearchInts(tails, n)
		if at == len(tails) {
			tails = append(tails, n)
		} else if tails[at] > n {
			tails[at] = n
		}
	}

	return len(tails)
}
