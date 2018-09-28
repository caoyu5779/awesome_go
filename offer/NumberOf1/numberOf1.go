package NumberOf1

func numberOf1(i int64) int {
	var count int
	var flag int64 = 1

	for flag != 0 {
		if i&flag != 0 {
			count += 1
		}
		flag <<= 1
	}

	return count
}
