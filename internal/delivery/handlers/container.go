package handlers

import (
	"github.com/gin-gonic/gin"
	"go-template/internal/service"
)

type ContainerHandler struct {
	service service.Container
}

func InitUserHandler(userService service.Container) ContainerHandler {
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
