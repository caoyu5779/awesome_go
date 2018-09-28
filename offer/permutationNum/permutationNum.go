package permutationNum

import (
	"fmt"
	"strconv"
	"strings"
)

func PermutationNum(n int) (resultSlice []string) {
	stringSlice := make([]string, n)
	permutationNum(stringSlice, 0, n-1, &resultSlice)
	return
}

func permutationNum(stringSlice []string, startIndex, endIndex int, resultSlice *[]string) {
	if startIndex == endIndex {
		for i := 0; i <= 9; i++ {
			stringSlice[endIndex] = strconv.Itoa(i)
			result := fmt.Sprint(strings.Join(stringSlice, ""))
			*resultSlice = append(*resultSlice, result)
		}
	} else {
		for i := startIndex; i <= endIndex; i++ {
			for num := 0; num <= 9; num++ {
				stringSlice[startIndex] = strconv.Itoa(num)
				permutationNum(stringSlice, startIndex+1, endIndex, resultSlice)
			}
		}
	}
}
