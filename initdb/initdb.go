package initdb

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID
	Username      string `gorm:"unique;not null"`
	Password      string `gorm:"not null"`
	Phone         string `gorm:"unique;not null"`
	CreatedAt     time.Time
	Conversations []Conversation
}

type Message struct {
	ID             uuid.UUID
	Status         string `gorm:"default:'unread'"`
	Body           string
	SendAt         time.Time
	ReadAt         time.Time
	UserID         string
	User           User
	ConversationID string
}

type Attachment struct {
	ID        uuid.UUID
	Type      string
	Url       string
	UploadAt  time.Time
	MessageID string
	Message   Message
}

type Conversation struct {
	ID       uuid.UUID
	UserID   string
	User     User
	Messages []Message
}

type Group struct {
	ID        uuid.UUID
	UserID    string
	User      User
	Name      string
	CreatedAt time.Time
}

type GroupMembers struct {
	ID      uuid.UUID
	GroupID string
	Group   Group
	UserID  string
	User    User
	Role    string `gorm:"default:'member'"`
}

func InitDb() {
	// connect to the postgres db using gorm
	dbName := "go_chat_db"
	connection := "host=localhost user=postgres password=postgres dbname=" + dbName + " port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// migration database schema
	if err := db.AutoMigrate(&Message{}, &Attachment{}, &Message{}, &User{}, &GroupMembers{}, &Group{}); err != nil {
		panic("failed to migrate database schema")
	}
	fmt.Println("Database initialized successfully")
}
