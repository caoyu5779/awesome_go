package maxProfit

func MaxProfit(price []int) int {
	length := len(price)
	res := 0

	for i := 1; i < length; i++ {
		if price[i] > price[i-1] {
			res += price[i] - price[i-1]
		}
	}

	return res
}
