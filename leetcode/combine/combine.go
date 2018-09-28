package combine

import "fmt"

func Combine(n int, k int) [][]int {
	combination := make([]int, k)
	res := [][]int{}

	var dfs func(int, int)
	dfs = func(idx int, begin int) {
		if idx == k {
			temp := make([]int, k)
			copy(temp, combination)
			fmt.Println(temp)
			res = append(res, temp)
			fmt.Println(res)
			fmt.Println()
			return
		}

		for i := begin; i <= n+1-k+idx; i++ {
			//idx, begin 0 1 k = 2
			//idx, begin 1,2 k =2
			//i = 1 i < 5-2+0=3 i++
			//i = 2 i < 5-2+1=4 i++
			combination[idx] = i
			//combination[0]=1
			//combination[1]=2
			//dfs(2,3)
			dfs(idx+1, i+1)
		}
	}

	dfs(0, 1)
	return res
}
