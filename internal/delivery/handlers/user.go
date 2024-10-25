package handlers

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/service"
)

type UserHandler struct {
	service service.User
}

func InitUserHandler(userService service.User) UserHandler {
	return UserHandler{
		service: userService,
	}
}

func (h UserHandler) Create(g *gin.Context) {

}

func (h UserHandler) GetUser(g *gin.Context) {

}

func (h UserHandler) Delete(g *gin.Context) {

}

func (h UserHandler) Login(g *gin.Context) {

}
