package main

import (
	"fmt"
	"golang-chat/maxprofit"
)

func main() {
	// get max profit from stock prices
	stokPrices := []int{3, 2, 6, 5, 1, 3} // sample price dari saham pada setiap harinya
	maxTransaction := 2                   // transaksi maksimal yang diperbolehkan

	maxprofit := maxprofit.MaxProfit(stokPrices, maxTransaction)

	// Menghitung dan mencetak keuntungan maksimal
	fmt.Println("Keuntungan maksimal yang dapat diperoleh:", maxprofit)

}
