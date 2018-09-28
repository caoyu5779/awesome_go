package cutTheRope

import (
	"math"
)

func cutTheRope(n int) int {
	if n < 2 {
		return 0
	}

	if n == 2 {
		return 1
	}

	if n == 3 {
		return 2
	}

	timesOfThree := n / 3

	if n-timesOfThree*3 == 1 {
		timesOfThree--
	}

	timesOfTwo := (n - timesOfThree*3) / 2
	return int(math.Pow(3, float64(timesOfThree)) * math.Pow(2, float64(timesOfTwo)))
}
