package validPalindrome

func ValidPalindrome(s string) bool {
	var bytes = []byte(s)

	low, high := 0, len(bytes) - 1
	return helper(bytes, low, high, false)
}

func helper(bs []byte, low, high int, hasDeleted bool) bool {
	for low < high {
		if bs[low] != bs[high] {
			if hasDeleted {
				return false
			}
			return helper(bs, low + 1, high, true) || helper(bs, low, high - 1, true)
		}
		low ++
		high --
	}

	return true
}
