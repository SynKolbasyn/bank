package domain

import (
	"errors"
	"net/http"

	"github.com/SynKolbasyn/bank/processor/internal/model"
	"github.com/labstack/echo/v5"
)

func ErrorResponse(ctx *echo.Context, err error) error {
	statusCode := http.StatusInternalServerError

	var appError *AppError
	if errors.As(err, &appError) {
		statusCode = appError.errCode
	}

	errorResponse := model.ErrorResponse{
		Error: http.StatusText(statusCode),
	}
	return ctx.JSON(appError.errCode, errorResponse)
}
