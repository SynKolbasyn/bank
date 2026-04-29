package app

import (
	echoprometheus "github.com/labstack/echo-prometheus"
	"github.com/labstack/echo/v5"
	echomw "github.com/labstack/echo/v5/middleware"
)

func setRoutes(server *echo.Echo, handlers *Handlers) {
	server.Use(echomw.Recover())
	server.Use(echoprometheus.NewMiddleware("payment-processing"))
	server.Use(echomw.RequestLogger())

	server.GET("/metrics", echoprometheus.NewHandler())
	server.GET("/health", handlers.health.Health)
}
