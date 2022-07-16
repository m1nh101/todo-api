package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m1ngi/todo-api/helper"
	"github.com/m1ngi/todo-api/model"
	"github.com/m1ngi/todo-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userAPI struct {
	repo repository.UserRepo
}

func NewUserAPI() *userAPI {
	return &userAPI{
		repo: repository.NewUserRepo(),
	}
}

func (u *userAPI) Login(c *gin.Context) {
	var credential model.User
	c.BindJSON(&credential)

	hashingPassword := helper.HashPassword(credential.Password)

	user, err := u.repo.FindUser(credential.Username, hashingPassword)

	if err == nil {
		token := helper.GenerateJwtToken(user)

		c.SetCookie("access_token", token, int(time.Now().Add(time.Hour*24).Unix()), "/", "localhost", true, true)
		c.SetSameSite(http.SameSiteNoneMode)
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"errors": err.Error()})
}

func (u *userAPI) Register(c *gin.Context) {
	var user model.User

	c.BindJSON(&user)

	user.Id = primitive.NewObjectID()

	hashingPassword := helper.HashPassword(user.Password)

	user.Password = hashingPassword

	response := u.repo.CreateUser(user)

	if response == nil {
		c.Status(http.StatusCreated)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"error": response.Error()})
}

func (u *userAPI) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", 0, "/", "", false, false)
	c.JSON(http.StatusNoContent, nil)
}
