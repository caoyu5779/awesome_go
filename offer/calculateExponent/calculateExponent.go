package calculateExponent

import (
	"errors"
	"math"
)

func calculateExponent(base float64, exponent int) (result float64, err error) {
	result = 1

	for i := 0; i < int(math.Abs(float64(exponent))); i++ {
		result *= base
	}

	if exponent < 0 {
		if base == 0 {
			err = errors.New("当exponent 为负数时， base 不能为0")
		}

		result = 1 / result
	}

	return
}
