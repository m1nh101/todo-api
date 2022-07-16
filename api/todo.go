package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m1ngi/todo-api/model"
	"github.com/m1ngi/todo-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type todoAPI struct {
	store repository.TodoRepo
}

func NewTodoAPI() *todoAPI {
	return &todoAPI{
		store: repository.NewTodoRepo(),
	}
}

func (to *todoAPI) CreateTodo(c *gin.Context) {
	var todo model.Todo

	userId := c.GetString("UserId")

	objectUserId, _ := primitive.ObjectIDFromHex(userId)

	c.BindJSON(&todo)

	todo.Id = primitive.NewObjectID()

	todo.UserId = objectUserId

	err := to.store.CreateTodo(todo)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

func (to *todoAPI) UpdateTodo(c *gin.Context) {
	var todo model.Todo

	c.BindJSON(&todo)

	todo.UpdateStatus()

	err := to.store.UpdateTodo(todo)

	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
	} else {
		c.JSON(http.StatusNoContent, todo)
	}
}

func (to *todoAPI) RemoveTodo(c *gin.Context) {
	todoId, succeed := c.Params.Get("id")
	userId := c.GetString("UserId")

	if !succeed {
		c.JSON(http.StatusBadRequest, gin.H{"errors": "id not found"})
		return
	}

	err := to.store.RemoveTodo(todoId, userId)

	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusAccepted, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (to *todoAPI) FetchTodo(c *gin.Context) {
	userId := c.GetString("UserId")

	response := to.store.GetTodo(userId)

	c.JSON(http.StatusNoContent, response)
}
