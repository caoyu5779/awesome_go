package findLongestWord

import "sort"

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] < s[j]
	}
	return len(s[i]) > len(s[j])
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isSub(s, sub string) bool {
	if len(s) < len(sub) {
		return false
	}

	i, j := 0, 0

	for i < len(s) && j < len(sub) {
		if s[i] == sub[j] {
			j++
		}
		i++
	}
	return j == len(sub)
}

func FindLongestWord(s string, d []string) string {
	sort.Sort(stringSlice(d))
	for i := range d {
		if isSub(s, d[i]) {
			return d[i]
		}
	}

	return ""
}
