package assignCookies

import "sort"

func AssignCookies(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var i,j,res int

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			res ++
			i ++
		}
		j ++
	}

	return res
}
