package chatmodel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	PhoneNumber string             `bson:"phone_number,omitempty"`
	Password    string             `bson:"password,omitempty"`
	CreatedAt   time.Time          `bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty"`
}

type Conversation struct {
	ID         primitive.ObjectID    `bson:"_id,omitempty"`
	SenderID   string    `bson:"sender_id,omitempty"`
	ReceiverID string    `bson:"receiver_id,omitempty"`
	Messages   []Message `bson:"messages"`
	CreatedAt  time.Time `bson:"created_at,omitempty"`
	UpdatedAt  time.Time `bson:"updated_at,omitempty"`
}

type Message struct {
	ID         primitive.ObjectID       `bson:"_id,omitempty"`
	SenderID   string       `bson:"sender_id,omitempty"`
	ReceiverID string       `bson:"receiver_id,omitempty"`
	Content    string       `bson:"content,omitempty"`
	Timestamp  time.Time    `bson:"timestamp,omitempty"`
	Status     string       `bson:"status,omitempty"`
	Attachment []Attachment `bson:"attachment"`
}

type Attachment struct {
	Name string `bson:"name,omitempty"`
	Type string `bson:"type,omitempty"`
	URL  string `bson:"url,omitempty"`
}

type Group struct {
	ID           primitive.ObjectID        `bson:"_id,omitempty"`
	Messages     []Message     `bson:"messages,omitempty"`
	CreatedBy    time.Time     `bson:"created_by,omitempty"`
	GroupMembers []GroupMember `bson:"group_members,omitempty"`
}

type GroupMember struct {
	UserId string `bson:"user_id,omitempty"`
	Role   string `bson:"role,omitempty"`
}
