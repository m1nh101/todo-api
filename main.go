package main

import (
	"github.com/gin-gonic/gin"
	"github.com/m1ngi/todo-api/api"
	"github.com/m1ngi/todo-api/middleware"
	"github.com/m1ngi/todo-api/repository"
)

func main() {
	repository.Setup()

	userAPI := api.NewUserAPI()

	todoAPI := api.NewTodoAPI()

	server := gin.Default()

	//auth route
	server.POST("api/auth/login", userAPI.Login)
	server.POST("api/auth/register", userAPI.Register)
	server.POST("api/auth/logout", userAPI.Logout)

	//todo doute
	server.GET("api/todos", middleware.AuthMiddleware(), todoAPI.FetchTodo)
	server.POST("api/todos", middleware.AuthMiddleware(), todoAPI.CreateTodo)
	server.PATCH("api/todos", middleware.AuthMiddleware(), todoAPI.UpdateTodo)
	server.DELETE("api/todos/:id", middleware.AuthMiddleware(), todoAPI.RemoveTodo)

	server.RunTLS(":2001", "./go-server.crt", "./go-server.key")
}
