package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"go-template/internal/delivery/handlers"
	"go-template/internal/delivery/handlers/middleware"
	"go-template/internal/repository/user"
	"go-template/internal/service"
	"go-template/pkg/logger"
)

func Start(db *sqlx.DB, logger logger.Logger) {
	r := gin.Default()
	r.ForwardedByClientIP = true

	containerRepo := user.InitUserRepo(db)
	containerService := service.InitContainerService(containerRepo, logger)
	containerHandler := handlers.InitUserHandler(containerService)

	userRouter := r.Group("/user")

	userRouter.POST("/create", containerHandler.Create)
	userRouter.GET("/get/:id", containerHandler.GetAll)
	userRouter.POST("/login", containerHandler.Login)

	mdw := middleware.InitMiddleware(&logger)
	r.Use(mdw.CORSMiddleware())

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
