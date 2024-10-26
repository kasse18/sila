package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"io"
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

func (h ContainerHandler) Upload(g *gin.Context) {
	containerID := g.Param("containerID")
	if containerID == "" {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Container ID is required"})
		return
	}

	file, _, err := g.Request.FormFile("file")
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}
	defer file.Close()

	url := "https://example-service-url.com/upload"

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header.Set("Content-Type", "multipart/form-data")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send request"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		g.JSON(http.StatusBadGateway, gin.H{"error": "Failed to upload file to external service"})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response"})
		return
	}

	id, ok := result["documentId"].(string)
	if !ok {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get documentId from response"})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message":    "File uploaded successfully",
		"documentId": id,
	})
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h ContainerHandler) Login(g *gin.Context) {
	var request LoginRequest
	login, password := "123", "123"
	if err := g.ShouldBindJSON(&request); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if request.Login != login || request.Password != password {
		g.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Login and password are incorrect")})
		return
	}
	secretKey := []byte("your_secret_key_here")

	claims := jwt.MapClaims{
		"login": login,
		"exp":   0,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	g.JSON(http.StatusOK, gin.H{
		"message": "Token generated successfully",
		"token":   tokenString,
	})
}
