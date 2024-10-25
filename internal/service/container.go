package service

import (
	"context"
	"go-template/internal/models/models"
	"go-template/internal/repository"
	"go-template/pkg/logger"
)

type containerService struct {
	containerRepo repository.ContainerRepo
	logger        logger.Logger
}

func InitContainerService(containerRepo repository.ContainerRepo, logger logger.Logger) User {
	return containerService{
		containerRepo: containerRepo,
		logger:        logger,
	}
}

func (c containerService) GetMe(ctx context.Context, id int) (*models.Container, error) {
	container, err := c.containerRepo.Get(ctx, id)
	if err != nil {
		c.logger.Error(ctx, err.Error())
		return nil, err
	}

	return container, nil
}

func (c containerService) Delete(ctx context.Context, id int) error {
	err := c.containerRepo.Delete(ctx, id)
	if err != nil {
		c.logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (c containerService) Create(ctx context.Context, container models.CreateContainer) (int, error) {
	return 0, nil
}

func (c containerService) Login(ctx context.Context, user models.CreateContainer) (int, error) {
	return 0, nil
}
