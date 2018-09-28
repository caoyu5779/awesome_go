package reconstructQueue

import "sort"

func ReconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0, len(people))
	sort.Sort(persons(res))

	insert := func(idx int, person []int) {
		res = append(res, person)
		if len(res)-1 == idx {
			return
		}

		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	for _, p := range people {
		insert(p[1], p)
	}

	return res
}

type persons [][]int

func (p persons) Len() int {
	return len(p)
}

func (p persons) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		return p[i][1] < p[j][1]
	}
	return p[i][0] > p[j][0]
}

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
