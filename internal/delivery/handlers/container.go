package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"sila-app/internal/models/models"
	"sila-app/internal/service"
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
	if err := g.ShouldBindJSON(&newContainer); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Вызываем сервис для создания контейнера
	err := h.service.Create(ctx, newContainer)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create container",
			"details": err.Error(),
		})
		return
	}

	// Отправляем успешный ответ
	g.JSON(http.StatusOK, gin.H{
		"message": "Container created successfully",
	})
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
