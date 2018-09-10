package findTopKFrequent

import "sort"

func FindTopKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)
	//
	rec := make(map[int]int , len(nums))

	for _, n := range nums {
		rec[n]++
	}

	counts := make([]int, len(rec))

	for _, c := range rec {
		counts = append(counts, c)
	}

	sort.Ints(counts)

	min := counts[len(counts) - k]

	for n,c := range rec {
		if c >= min {
			res = append(res, n)
		}
	}

	return res
}
