package myAtoi

import (
	"strings"
)

const (
	MAX int32 = 1 << 31 -1
	MIN int32 = -1 * MAX - 1
)

func MyAtoi(str string) int32 {
	// 除去 str 首尾的空格
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	// 取出 s 的符号和主体 x
	var sign int32
	var x string
	sign, x = getSign(s)

	// 裁剪x丢弃混入的非数字字符
	x = trim(x)

	// 根据sigh和x，转换成int
	return convert(sign, x)
}

func getSign(s string) (int32, string) {
	var sign int32
	sign = 1
	switch s[0] {
	case '-':
		s = s[1:]
		sign = -1.0
	case '+':
		s = s[1:]
	default:
	}

	return int32(sign), s
}

func trim(s string) string {
	for i := range s {
		if s[i] < '0' || '9' < s[i] {
			return s[:i]
		}
	}

	return s
}

func convert(sign int32, s string) int32 {
	var base int32
	base = 1 * sign
	var res int32
	res = 0
	yes := false

	for i := len(s) - 1; i >= 0; i-- {
		// 为了防止特别长的数字，甚至超过float64的范围，所以，每一步都检查是否溢出
		res, yes = isOverflow(int32(res + (int32(s[i])-48)*base))
		if yes {
			return res
		}
		base *= 10
	}

	return int32(res)
}

func isOverflow(i int32) (int32, bool) {
	switch {
	case i > MAX:
		return MAX, true
	case i < MIN:
		return MIN, true
	default:
		return i, false
	}
}

