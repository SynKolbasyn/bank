package handler

import (
	"net/http"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/SynKolbasyn/bank/gateway/internal/service"
	"github.com/labstack/echo/v5"
)

type Auth struct {
	serviceAuth service.IAuth
}

func NewAuth(serviceAuth service.IAuth) *Auth {
	return &Auth{
		serviceAuth: serviceAuth,
	}
}

func (a *Auth) SignUp(ctx *echo.Context) error {
	var signRequest model.SignRequest

	err := ctx.Bind(&signRequest)
	if err != nil {
		return domain.ErrorResponse(
			ctx,
			domain.NewAppError(http.StatusBadRequest, err),
		)
	}

	err = ctx.Validate(signRequest)
	if err != nil {
		return domain.ErrorResponse(
			ctx,
			domain.NewAppError(http.StatusBadRequest, err),
		)
	}

	token, err := a.serviceAuth.SignUp(ctx.Request().Context(), &signRequest)
	if err != nil {
		return domain.ErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusCreated, model.TokenResponse{
		Token: token,
	})
}

func (a *Auth) SignIn(ctx *echo.Context) error {
	var signRequest model.SignRequest

	err := ctx.Bind(&signRequest)
	if err != nil {
		return domain.ErrorResponse(
			ctx,
			domain.NewAppError(http.StatusBadRequest, err),
		)
	}

	err = ctx.Validate(signRequest)
	if err != nil {
		return domain.ErrorResponse(
			ctx,
			domain.NewAppError(http.StatusBadRequest, err),
		)
	}

	token, err := a.serviceAuth.SignIn(ctx.Request().Context(), &signRequest)
	if err != nil {
		return domain.ErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, model.TokenResponse{
		Token: token,
	})
}
