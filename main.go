package main

import (
	"fmt"
	"golang-chat/initdb"
	"golang-chat/maxprofit"
)

func main() {

	// create database schema for basic chat messaging system with postgresql
	initdb.InitDb()

	// get max profit from stock prices
	stokPrices := []int{3, 2, 6, 5, 1, 3} // sample price stok list of each day
	maxTransaction := 2 // maximum transaction allowed

	maxprofit := maxprofit.MaxProfit(stokPrices, maxTransaction)

	// Menghitung dan mencetak keuntungan maksimal
	fmt.Println("Keuntungan maksimal yang dapat diperoleh:", maxprofit)

}
