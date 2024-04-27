package mongoactions

import (
	"context"
	"fmt"
	"golang-chat/chatmodel"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://localhost:27017"

var (
	client *mongo.Client
	db     *mongo.Database
	ctx    context.Context
)

func Connect() error {
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	db = client.Database("chat")
	ctx = context.Background()
	fmt.Println("Successfully connected to MongoDB")
	return nil
}

func Disconnect() {
	client.Disconnect(ctx)
}

func CreateUser(name, phone_number, password string) (*mongo.InsertOneResult, error) {
	user := chatmodel.User{
		ID:          primitive.NewObjectID(),
		Name:        name,
		PhoneNumber: phone_number,
		Password:    password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result, err := db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateMessage(senderID, receiverID, content, status string, attachment []chatmodel.Attachment) chatmodel.Message {
	message := chatmodel.Message{
		ID:         primitive.NewObjectID(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		Timestamp:  time.Now(),
		Status:     status,
		Attachment: attachment,
	}

	return message
}

func CreateAttachment(name, tipe, url string) chatmodel.Attachment {
	attachment := chatmodel.Attachment{
		Name: name,
		Type: tipe,
		URL:  url,
	}

	return attachment
}

func CreateConversation(senderID, receiverID string, messages []chatmodel.Message) (*mongo.InsertOneResult, error) {
	conversation := chatmodel.Conversation{
		ID:         primitive.NewObjectID(),
		SenderID:   senderID,
		ReceiverID: receiverID,
		Messages:   messages,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	result, err := db.Collection("conversations").InsertOne(ctx, conversation)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateGroup(createdBy string, messages []chatmodel.Message, groupMembers []chatmodel.GroupMember) (*mongo.InsertOneResult, error) {
	group := chatmodel.Group{
		ID:           primitive.NewObjectID(),
		Messages:     messages,
		CreatedBy:    time.Now(),
		GroupMembers: groupMembers,
	}

	result, err := db.Collection("groups").InsertOne(ctx, group)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func CreateGroupMember(userID, role string) chatmodel.GroupMember {
	groupMember := chatmodel.GroupMember{
		UserId: userID,
		Role:   role,
	}

	return groupMember
}
