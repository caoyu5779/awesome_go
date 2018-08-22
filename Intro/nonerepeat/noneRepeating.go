package main

import "fmt"

//var lastOccurred = make([]int, 0xffff)

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	//for i := range lastOccurred {
	//	lastOccurred[i] = 0
	//}

	for i, ch := range []rune(s) {

		if lastI, ok:= lastOccurred[ch] ; ok && lastI > start {
			start = lastI + 1
		}

		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("我爱拜仁"))
	// strings fields split join
	// strings contains index
	// strings ToLower ToUpper
	// strings Trim TrimRight TrimLeft
	// utf8 转乘rune操作
}
