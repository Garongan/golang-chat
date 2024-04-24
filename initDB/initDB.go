package initDB

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Username  string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Phone     string    `gorm:"unique;not null"`
	CreatedAt time.Time
}

type Messages struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	SenderID   Users     `gorm:"foreignKey:ID"`
	ReceiverID Users     `gorm:"foreignKey:ID"`
	Status     string    `gorm:"default:'unread'"`
	Message    string
	SendAt     time.Time
	ReadAt     time.Time
}

type Attachments struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	MessageID Messages  `gorm:"foreignKey:ID"`
	Type      string
	Url       string
	UploadAt  time.Time
}

type Groups struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	CreatedBy Users     `gorm:"foreignKey:ID"`
	Name      string
	CreatedAt time.Time
}

type GroupMembers struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;"`
	GroupID Groups    `gorm:"foreignKey:ID"`
	UserID  Users     `gorm:"foreignKey:ID"`
	Role    string    `gorm:"default:'member'"`
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
	if err := db.AutoMigrate(&Attachments{}, &Messages{}, &GroupMembers{}, &Groups{}, &Users{}); err != nil {
		panic("failed to migrate database schema")
	}
	fmt.Println("Database initialized successfully")
}
