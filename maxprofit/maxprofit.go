package maxprofit

import (
	"fmt"
	"math"
)

func MaxProfit(prices []int, maxTransaction int) int {
	if maxTransaction >= len(prices)/2 {
		hold := -prices[0]
		sell := 0

		for i := 0; i < len(prices); i++ {
			hold = max(hold, sell-prices[i])
			sell = max(sell, hold+prices[i])
			fmt.Println(sell)
		}

		return sell
	}

	matriksProfit := make([][]int, maxTransaction+1)
	for i := range matriksProfit {
		matriksProfit[i] = make([]int, len(prices))
	}

	for i := 1; i <= maxTransaction; i++ {
		hold := math.MinInt32
		for j := 1; j < len(prices); j++ {
			hold = max(hold, matriksProfit[i-1][j-1] - prices[j-1])
			matriksProfit[i][j] = max(matriksProfit[i][j-1], hold + prices[j])
		}
	}

	return matriksProfit[maxTransaction][len(prices)-1]
}
