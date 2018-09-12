package findMin

func FindMin(num []int) int {
	Len := len(num)

	i := 1

	for i < Len && num[i-1] < num[i] {
		i++
	}
	return num[i%Len]
}
