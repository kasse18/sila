package middleware

import "go-template/pkg/logger"

type Middleware struct {
	logger *logger.Logger
}

func InitMiddleware(logger *logger.Logger) Middleware {
	return Middleware{
		logger: logger,
	}
}
