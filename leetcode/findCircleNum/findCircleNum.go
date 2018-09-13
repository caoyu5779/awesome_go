package findCircleNum

func FindCircleNum(M [][]int) int {
	N := len(M)

	res := N

	friend := make([]int, res)
	for i := 0; i < res ; i ++ {
		friend[i] = i
 	}

 	union := func(s, d int) {
 		for i := range friend {
 			if friend[i] == s{
 				friend[i] = d
			}
		}
	}


	for i := 0; i < N; i ++ {
		for j := i+1; j < N; j++ {
			if M[i][j] == 1 {
				if friend[i] != friend[j] {
					res --
					union(friend[i], friend[j])
				}
			}
		}
	}

	return res
}
