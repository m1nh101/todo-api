package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      primitive.ObjectID `bson:"user_id"`
	Content     string             `bson:"content" json:"content"`
	CreatedAt   time.Time          `bson:"create_at"`
	Priority    int8               `bson:"priority" json:"priority"`
	Completed   bool               `bson:"completed" json:"completed"`
	CompletedAt time.Time          `bson:"completed_at"`
}

func (todo *Todo) UpdateStatus() {
	todo.Completed = !todo.Completed
	todo.CompletedAt = time.Now().UTC()
}
