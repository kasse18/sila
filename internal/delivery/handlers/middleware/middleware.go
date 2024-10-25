package middleware

import "sila-app/pkg/logger"

type Middleware struct {
	logger *logger.Logger
}

func InitMiddleware(logger *logger.Logger) Middleware {
	return Middleware{
		logger: logger,
	}
}
