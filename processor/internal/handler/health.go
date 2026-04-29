package handler

import (
	"net/http"

	"github.com/SynKolbasyn/bank/processor/internal/domain"
	"github.com/SynKolbasyn/bank/processor/internal/service"
	"github.com/labstack/echo/v5"
)

type Health struct {
	serviceHealth service.IHealth
}

func NewHealth(serviceHealth service.IHealth) *Health {
	return &Health{
		serviceHealth: serviceHealth,
	}
}

func (h *Health) Health(ctx *echo.Context) error {
	err := h.serviceHealth.Health(ctx.Request().Context())
	if err != nil {
		return domain.ErrorResponse(ctx, err)
	}

	return ctx.NoContent(http.StatusOK)
}
