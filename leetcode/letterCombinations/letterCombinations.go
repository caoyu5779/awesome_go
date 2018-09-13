package letterCombinations

var m = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	ret := []string{""}

	for i:=0; i < len(digits); i++ {
		temp := []string{}

		for j:=0; j < len(ret); j ++ {
			for k :=0; k < len(m[digits[i]]); k ++{
				temp = append(temp, ret[j]+m[digits[i]][k])
			}
		}

		ret = temp
	}

	return ret
}
