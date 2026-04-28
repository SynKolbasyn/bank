package handler

import (
	"net/http"

	"github.com/SynKolbasyn/bank/internal/service"
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
	health := h.serviceHealth.Health(ctx.Request().Context())
	return ctx.JSON(http.StatusOK, health)
}
