package partitionLabels

func PartitionLabels(S string) []int  {
	maxIndex :=[26]int{}

	for i,b := range S {
		maxIndex[b-'a'] = i
	}

	begin := 0
	end := maxIndex[S[begin]-'a']

	res := make([]int, 0, len(S))

	for i,b := range S {
		if i < end {
			end = max (end,maxIndex[b-'a'])
			continue
		}
		res = append(res, i-begin+1)
		begin = i + 1
		if begin < len(S) {
			end = maxIndex[S[begin] - 'a']
		}
	}

	return res
}

func max(a,b int) int {
	if a > b {
		return a
	}

	return b
}