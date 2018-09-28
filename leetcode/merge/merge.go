package merge

func Merge(num1 []int, m int, num2 []int, n int) []int {
	temp := make([]int, m)
	copy(temp, num1)

	j, k := 0, 0

	for i := 0; i < len(num1); i++ {
		if k >= n {
			num1[i] = temp[j]
			j++
			continue
		}

		if j >= m {
			num1[i] = num2[k]
			k++
			continue
		}

		if temp[j] < num2[k] {
			num1[i] = temp[j]
			j++
		} else {
			num1[i] = num2[k]
			k++
		}
	}

	return num1
}
