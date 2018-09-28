package reverseVowels

func ReverseVowels(s string) string {
	//只遍历一次
	bytes := []byte(s)

	i, j := 0, len(bytes)-1

	for {
		for i < len(s) && !isVowels(bytes[i]) {
			i++
		}

		for 0 < j && !isVowels(bytes[j]) {
			j--
		}

		if i >= j {
			break
		}

		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	return string(bytes)
}

func isVowels(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
