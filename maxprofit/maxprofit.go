package maxprofit

import (
	"fmt"
	"math"
)

func MaxProfit(prices []int, maxTransaction int) int {
	// kondisi ketika maxTransaction lebih besar atau sama dengan setengat dari panjang prices
	// dilakukan seluruh transaksi setiap hari
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

	// membuat matriks untuk menyimpan profit
	// setiap baris menunjukkan transaksi ke-i, dengan panjangnya adalah maxtransaction+1
	matriksProfit := make([][]int, maxTransaction+1)
	for i := range matriksProfit {
		// menginisialisasi matriks dengan panjang prices
		// digunakan untuk menyimpan kemungkinan profit yang didapat
		matriksProfit[i] = make([]int, len(prices))
	}

	for i := 1; i <= maxTransaction; i++ {
		// mencari selisih profit dari transaksi ke-i
		different := math.MinInt32
		for j := 1; j < len(prices); j++ {
			different = max(different, matriksProfit[i-1][j-1] - prices[j-1])
			matriksProfit[i][j] = max(matriksProfit[i][j-1], different + prices[j])
		}
	}

	// mengambalikan profit terbanyak yang didapat dari indeks terakhir di baris terakhir pada matrik profit
	return matriksProfit[maxTransaction][len(prices)-1]
}
