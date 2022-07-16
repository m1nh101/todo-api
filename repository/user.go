package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/m1ngi/todo-api/model"
)

type UserRepo interface {
	CreateUser(user model.User) error
	FindUser(username, password string) (*model.User, error)
}

type userRepo struct{}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (u *userRepo) CreateUser(user model.User) error {

	result, err := UserCollection.InsertOne(context.TODO(), user)

	fmt.Print(result.InsertedID)

	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (u *userRepo) FindUser(username, password string) (*model.User, error) {

	var user model.User

	UserCollection.FindOne(context.TODO(), model.User{Username: username, Password: password}).Decode(&user)

	if user.Id.String() == "" {
		return nil, errors.New("wrong credential")
	}

	return &user, nil
}
