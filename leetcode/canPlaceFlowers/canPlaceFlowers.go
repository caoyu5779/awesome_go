package canPlaceFlowers

func CanPlaceFlowers(flowered []int, n int) bool  {
	l := len(flowered)

	for i:=0; i < l; i++ {
		if flowered[i] == 0 &&
			((i + 1 < l) && flowered[i+1] == 0 || i + 1 >= l) &&
			((i-1 >= 0) && flowered[i-1] == 0 || i - 1 < 0) {
				flowered[i] = 1
				n --
				if n <= 0 {
					return true
				}
		}
	}

	return n <= 0
}
