package mySqrt

func MySqrt(x int) int {
	res := x

	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}
