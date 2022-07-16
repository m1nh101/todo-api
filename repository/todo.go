package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/m1ngi/todo-api/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepo interface {
	GetTodo(userId string) []model.Todo
	CreateTodo(todo model.Todo) error
	UpdateTodo(todo model.Todo) error
	RemoveTodo(id, userId string) error
}

type todoRepo struct{}

func NewTodoRepo() *todoRepo {
	return &todoRepo{}
}

func (re *todoRepo) GetTodo(userId string) []model.Todo {

	var todos []model.Todo

	objectUserId, _ := primitive.ObjectIDFromHex(userId)

	cursor, _ := TodoCollection.Find(context.TODO(), bson.M{"user_id": objectUserId})

	for cursor.Next(context.TODO()) {
		var todo model.Todo

		cursor.Decode(&todo)

		todos = append(todos, todo)
	}

	return todos
}

func (re *todoRepo) CreateTodo(todo model.Todo) error {

	todo.CreatedAt = time.Now().UTC()
	todo.Completed = false
	todo.CompletedAt = time.Now().UTC()

	result, err := TodoCollection.InsertOne(context.TODO(), todo)

	fmt.Print(result.InsertedID)

	return err
}

func (re *todoRepo) UpdateTodo(todo model.Todo) error {

	context, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	defer cancel()

	result, err := TodoCollection.UpdateByID(context, todo.Id, bson.M{"$set": todo})

	fmt.Print(result.UpsertedID)

	return err
}

func (re *todoRepo) RemoveTodo(id, userId string) error {

	objectID, _ := primitive.ObjectIDFromHex(id)

	userObjectId, _ := primitive.ObjectIDFromHex(userId)

	result, err := TodoCollection.DeleteOne(context.TODO(), model.Todo{Id: objectID, UserId: userObjectId})

	fmt.Print(result.DeletedCount)

	return err
}
