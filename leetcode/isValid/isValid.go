package isValid

type stack []rune

func (s *stack) Push(b rune) {
	*s = append(*s, b)
}

func (s *stack) Pop() (rune, bool) {
	if len(*s) > 0 {
		res := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return res, true
	}
	return 0, false
}

var matching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func IsValid(str string) bool {
	s := new(stack)

	for _, b := range str {
		switch b {
		case '(', '{', '[':
			s.Push(b)
		case ')', '}', ']':
			if r, ok := s.Pop(); !ok || r != matching[b] {
				return false
			}
		}
	}

	if len(*s) > 0 {
		return false
	}

	return true
}
