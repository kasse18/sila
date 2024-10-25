package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-template/internal/models/models"
	"go-template/internal/service"
	"net/http"
	"time"
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
	var newContainer models.CreateContainer

	if err := g.ShouldBind(&newContainer); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := h.service.Create(ctx, newContainer)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"id": id})
}

func (h ContainerHandler) GetAll(g *gin.Context) {
	ctx := g.Request.Context()

	containers, err := h.service.GetAll(ctx)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g.JSON(http.StatusOK, gin.H{"containers": containers})
}

func (h ContainerHandler) Login(g *gin.Context) {

}

func (h ContainerHandler) Upload(g *gin.Context) {

}
