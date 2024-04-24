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
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2

	maxprofit := maxprofit.MaxProfit(prices, k)

	// Menghitung dan mencetak keuntungan maksimal
	fmt.Println("Keuntungan maksimal yang dapat diperoleh:", maxprofit)

}
