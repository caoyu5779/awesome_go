package nextGreatestLetter

import "sort"

func NextGreatestLetter(letters []byte, target byte) byte {
	n := len(letters)

	i := sort.Search(n, func(i int) bool {
		return target < letters[i]
	})

	return letters[i]
}
