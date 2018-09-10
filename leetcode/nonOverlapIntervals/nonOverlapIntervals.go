package nonOverlapIntervals

import (
	"selfLearning/leetcode/tool"
	"sort"
)

type Interval = tool.Interval

func NonOverlapIntervals(intervals []Interval) int {
	end := -1 << 31
	res := 0

	if len(intervals) == 1 {
		return 0
	}

	sort.Sort(B(intervals))

	for _, v := range intervals {
		if end <= v.Start{
			end = v.End
		} else {
			res ++
		}
	}
	return res
}

type B []Interval

func (a B) Len() int {
	return len(a)
}

func (a B) Swap(i,j int) {
	a[i], a[j] = a[j], a[i]
}

func (a B) Less(i,j int) bool  {
	return a[i].End < a[j].End
}
