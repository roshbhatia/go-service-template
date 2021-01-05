package api

import (
	"context"

	"github.com/roshbhatia/echo-service/logger"
)

type Api struct {
	Logger logger.Logger
	Ctx    context.Context
}
