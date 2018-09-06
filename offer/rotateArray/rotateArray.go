package rotateArray

func rotateArray(numbers []int, k int) []int {
	arrayLength := len(numbers)

	if arrayLength == 0{
		return numbers
	}

	numbers = append(numbers[k+1 :], numbers[:k + 1]...)
	return numbers
}

