package dailyTemperatures

import "fmt"

func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	res := make([]int, n)

	stack := make([]int, n)

	top := -1

	for i := 0; i < n; i++ {
		for top >= 0 && temperatures[stack[top]] < temperatures[i] {
			res[stack[top]] = i - stack[top]
			fmt.Println(res)
			top--
			fmt.Println(top)
		}
		top++
		fmt.Println(top)
		fmt.Println()
		stack[top] = i
	}

	return res
}
