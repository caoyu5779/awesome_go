package rotateArray

func FindMin(array []int) int {
	return findMin(array, 0, len(array)-1)
}

func findMin(array []int, low, high int) int {
	if low > high {
		return array[0]
	}

	middle := (low + high) / 2

	if array[middle] > array[middle+1] {
		return array[middle+1]
	}

	if array[middle] > array[high] {
		return findMin(array, middle+1, high)
	} else {
		return findMin(array, low, middle-1)
	}
}
