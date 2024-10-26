package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"sila-app/internal/delivery/handlers"
	"sila-app/internal/delivery/handlers/middleware"
	"sila-app/internal/repository/container"
	"sila-app/internal/service/container"
	"sila-app/pkg/logger"
)

func Start(db *sqlx.DB, logger *logger.Logger) {
	r := gin.Default()
	r.ForwardedByClientIP = true

	mdw := middleware.InitMiddleware(logger)
	r.Use(mdw.CORSMiddleware())

	containerRepo := repository.InitContainerRepo(db, logger)
	containerService := service.InitContainerService(containerRepo, logger)
	containerHandler := handlers.InitUserHandler(containerService)

	userRouter := r.Group("/container")

	userRouter.POST("/create_container", containerHandler.Create)
	userRouter.GET("/get_all_containers", containerHandler.GetAll)
	userRouter.POST("/login", containerHandler.Login)
	userRouter.POST("/upload/:containerID", containerHandler.Upload)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
