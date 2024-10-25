package handlers

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/service"
)

type ContainerHandler struct {
	service service.User
}

func InitUserHandler(userService service.User) ContainerHandler {
	return ContainerHandler{
		service: userService,
	}
}

func (h ContainerHandler) Create(g *gin.Context) {

}

func (h ContainerHandler) GetAll(g *gin.Context) {

}

func (h ContainerHandler) Login(g *gin.Context) {

}

func (h ContainerHandler) Upload(g *gin.Context) {

}
