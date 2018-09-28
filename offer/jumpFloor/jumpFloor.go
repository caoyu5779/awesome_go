package jumpFloor

func JumpFloor(i int) int {
	if i == 1 || i == 2 {
		return i
	}

	return JumpFloor(i-2) + JumpFloor(i-1)
}
