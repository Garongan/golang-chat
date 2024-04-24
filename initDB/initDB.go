package initDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type Message struct {
	ID         string `json:"id"`
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
	TimeStamp  string `json:"timestamp"`
}

type Group struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GroupMember struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

type File struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Filename  string `json:"filename"`
	Filtetype string `json:"filetype"`
	Size      string `json:"size"`
	SharedAt  string `json:"shared_at"`
}

func InitDb() {
	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=go_chat_db user=postgres password=postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create user table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(50) NOT NULL,
		phone VARCHAR(15) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	// create message table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		sender_id INT NOT NULL,
		receiver_id INT NOT NULL,
		content TEXT NOT NULL,
		timestamp TIMESTAMP NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	// create group table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS groups (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	// create group_member table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS group_members (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		group_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	// create file table
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS files (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		filename VARCHAR(50) NOT NULL,
		filetype VARCHAR(50) NOT NULL,
		size VARCHAR(50) NOT NULL,
		shared_at TIMESTAMP NOT NULL
	)`)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database initialized successfully")
}
