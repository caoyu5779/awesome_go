package checkPossibility

import (
	"sort"
)

func CheckPossibility(num []int) bool {
	for i := 1; i < len(num); i++ {
		if num[i-1] > num[i] {
			pre := deepCopy(num)

			pre[i-1] = pre[i]

			next := deepCopy(num)
			next[i] = next[i-1]

			return sort.IsSorted(sort.IntSlice(pre)) || sort.IsSorted(sort.IntSlice(next))
		}
	}

	return true
}

func deepCopy(num []int) []int {
	res := make([]int, len(num))

	copy(res, num)

	return res
}
