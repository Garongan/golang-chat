package main

import (
	"fmt"
	"golang-chat/chatmodel"
	"golang-chat/maxprofit"
	"golang-chat/mongoactions"
)

func main() {

	// Connect to mongodb localhost
	err := mongoactions.Connect()
	if err != nil {
		fmt.Println(err)
	}

	defer mongoactions.Disconnect()

	// create new user
	/*
		newUser, err := mongoactions.CreateUser("John Babi", "08123456789", "password")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("new user: ", newUser)
	*/

	// create new conversations
	// create new list of messages
	/*
		newMessage := mongoactions.CreateMessage("sender_id", "receiver_id", "Hello, how are you?", "unread", nil)

		newMessages := make([]chatmodel.Message, 0)
		newMessages = append(newMessages, newMessage)

		newConversation, err := mongoactions.CreateConversation("sender_id", "receiver_id", newMessages)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("new conversation: ", newConversation)
	*/

	// create new list of message
	// create new group
	/*
		newMessage := mongoactions.CreateMessage("sender_id", "receiver_id", "Hello, how are you?", "unread", nil)
		newMessages := make([]chatmodel.Message, 0)
		newMessages = append(newMessages, newMessage)

		newGroupMembers := make([]chatmodel.GroupMember, 0)
		newGroupMembers = append(newGroupMembers, mongoactions.CreateGroupMember("admin_id", "admin"))

		newGroup, err := mongoactions.CreateGroup("admin_id", newMessages, newGroupMembers)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("new group: ", newGroup)
	*/
	
	// get max profit from stock prices
	stokPrices := []int{3, 2, 6, 5, 1, 3} // sample price dari saham pada setiap harinya
	maxTransaction := 2                   // transaksi maksimal yang diperbolehkan

	maxprofit := maxprofit.MaxProfit(stokPrices, maxTransaction)

	// Menghitung dan mencetak keuntungan maksimal
	fmt.Println("Keuntungan maksimal yang dapat diperoleh:", maxprofit)
}
