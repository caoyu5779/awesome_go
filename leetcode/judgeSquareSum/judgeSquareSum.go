package judgeSquareSum

import "math"

func JudgeSquareSum(number int) bool {
	a := intSqrt(number)

	for a >= 0 {
		if isSquare(number - a*a) {
			return true
		}
	}

	return false

}

func intSqrt(c int) int {
	return int(math.Sqrt(float64(c)))
}

func isSquare(a int) bool {
	b := intSqrt(a)
	return a*a == b
}
