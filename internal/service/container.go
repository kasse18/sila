package service

import (
	"context"
	"sila-app/internal/models/models"
	"sila-app/internal/repository"
	"sila-app/pkg/logger"
)

type containerService struct {
	containerRepo repository.ContainerRepo
	logger        logger.Logger
}

func InitContainerService(containerRepo repository.ContainerRepo, logger *logger.Logger) Container {
	return containerService{
		containerRepo: containerRepo,
		logger:        *logger,
	}
}

func (c containerService) Upload(ctx context.Context) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c containerService) GetAll(ctx context.Context) ([]models.Container, error) {
	container, err := c.containerRepo.GetAll(ctx)
	if err != nil {
		c.logger.Error(ctx, err.Error())
		return nil, err
	}

	return container, nil
}

func (c containerService) Create(ctx context.Context, container models.CreateContainer) error {
	err := c.containerRepo.Create(ctx, container)
	if err != nil {
		c.logger.Error(ctx, err.Error())
		return err
	}

	return nil
}

func (c containerService) Login(ctx context.Context, user models.CreateContainer) (int, error) {
	return 0, nil
}
