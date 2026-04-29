package domain

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/labstack/echo/v5"
)

const (
	minClientErrCode = http.StatusBadRequest
	minServerErrCode = http.StatusInternalServerError
)

func ErrorResponse(ctx *echo.Context, err error) error {
	statusCode := http.StatusInternalServerError

	if err != nil {
		appError, ok := errors.AsType[*AppError](err)
		if ok {
			switch {
			case minClientErrCode <= appError.statusCode && appError.statusCode < minServerErrCode:
				ctx.Logger().WarnContext(
					ctx.Request().Context(),
					"client error",
					slog.Int("status_code", appError.statusCode),
					slog.String("error", appError.Error()),
				)
			case minServerErrCode <= appError.statusCode:
				ctx.Logger().ErrorContext(
					ctx.Request().Context(),
					"server error",
					slog.Int("status_code", appError.statusCode),
					slog.String("error", appError.Error()),
				)
			default:
				ctx.Logger().WarnContext(
					ctx.Request().Context(),
					"calling `domain.ErrorResponse` with non error status code",
					slog.Int("status_code", appError.statusCode),
					slog.String("error", appError.Error()),
				)
			}

			statusCode = appError.statusCode
		} else {
			ctx.Logger().
				ErrorContext(ctx.Request().Context(), "unknown server error", slog.String("error", err.Error()))
		}
	} else {
		ctx.Logger().
			WarnContext(ctx.Request().Context(), "calling `domain.ErrorResponse` with nil error")
	}

	errorResponse := model.ErrorResponse{
		Error: http.StatusText(statusCode),
	}

	return ctx.JSON(statusCode, errorResponse)
}
