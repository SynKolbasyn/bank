package app

import (
	"github.com/SynKolbasyn/bank/config"
	"github.com/SynKolbasyn/bank/internal/middleware"
	echojwt "github.com/labstack/echo-jwt/v5"
	echoprometheus "github.com/labstack/echo-prometheus"
	"github.com/labstack/echo/v5"
	echomw "github.com/labstack/echo/v5/middleware"
)

func setRoutes(server *echo.Echo, config *config.Config, handlers *Handlers) {
	jwtConfig := middleware.NewJWTConfig(config.Auth.Secret)

	server.Use(echomw.Recover())
	server.Use(echoprometheus.NewMiddleware("api-gateway"))
	server.Use(echomw.RequestLogger())
	server.Use(echomw.CORS("http://localhost:5173", "http://localhost:4173", "http://localhost:80"))

	server.GET("/metrics", echoprometheus.NewHandler())
	server.GET("/health", handlers.health.Health)

	server.POST("/auth/sign-up", handlers.auth.SignUp)
	server.POST("/auth/sign-in", handlers.auth.SignIn)

	server.POST("/payments", handlers.payments.Create, echojwt.WithConfig(jwtConfig))
	server.GET("/payments", handlers.payments.Get, echojwt.WithConfig(jwtConfig))
}
