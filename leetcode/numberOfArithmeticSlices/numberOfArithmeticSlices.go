package numberOfArithmeticSlices

import "fmt"

func NumberOfArithmeticSlice(a []int) int {
	if len(a) < 3 {
		return 0
	}

	res := 0

	var i,j = 0,0

	for i < len(a) {
		j = i + 2
		//2
		//5
		for j < len(a) && a[j] - a[j-1] == a[j-1] - a [j-2] {
			j ++
			//3
			//4
		}
		j--
		//3
		//3
		fmt.Println(j)
		res += (j-i-1) * (j - i) / 2
		i = j
		//3
	}
	return res
}
