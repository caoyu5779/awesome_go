package findLongestChain

import "sort"

func FindLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i][1] == pairs[j][1] {
			return pairs[i][0] > pairs[j][0]
		}

		return pairs[i][1] < pairs[j][1]
	})
	//先进行了排序
	res := 0

	b := -1 << 32

	for i := 0; i < len(pairs); i++ {
		c := pairs[i][0]
		if b < c {
			res++
			b = pairs[i][1]
		}
	}

	return res
}
