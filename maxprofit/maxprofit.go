package maxprofit

func MaxProfit(prices []int, k int) int {
	n := len(prices)
	// Kasus k terlalu besar, dapat melakukan transaksi sebanyak yang diinginkan
	if k >= n/2 {
		maxProfit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				maxProfit += prices[i] - prices[i-1]
			}
		}
		return maxProfit
	}

	dp := make([][]int, k+1)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i <= k; i++ {
		maxPrevDiff := -prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+maxPrevDiff)
			maxPrevDiff = max(maxPrevDiff, dp[i-1][j]-prices[j])
		}
	}

	return dp[k][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}