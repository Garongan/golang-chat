package maxprofit

func MaxProfit(prices []int, maxTransaction int) int {
	profit := 0
	if len(prices) == 1 {
		// tidak ada keuntungan
		return profit
	}

	// pembelian saham pertama
	buy := prices[0]
	for index, price := range prices {
		transaction := 0
		if transaction < maxTransaction && index+1 < len(prices) && buy != 0 && buy < price {
			// penjualan => profit bertambah
			profitTemp := price - buy
			transaction++

			if profitTemp > profit {
				profit = profitTemp
			}

			if profit > price {
				buy = prices[index+1]
			} else {
				buy = 0
			}
		}
	}

	return profit
}
