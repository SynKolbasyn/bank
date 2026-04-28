package handler

import (
	"net/http"

	"github.com/SynKolbasyn/bank/internal/domain"
	"github.com/SynKolbasyn/bank/internal/middleware"
	"github.com/SynKolbasyn/bank/internal/model"
	"github.com/SynKolbasyn/bank/internal/service"
	"github.com/labstack/echo/v5"
)

type Payments struct {
	servicePayments service.IPayments
}

func NewPayments(servicePayments service.IPayments) *Payments {
	return &Payments{
		servicePayments: servicePayments,
	}
}

func (p *Payments) Create(ctx *echo.Context) error {
	var paymentRequest model.PaymentRequest
	err := ctx.Bind(&paymentRequest)
	if err != nil {
		return domain.ErrorResponse(
			ctx,
			domain.NewAppError(http.StatusBadRequest, err),
		)
	}
	err = ctx.Validate(paymentRequest)
	if err != nil {
		return domain.ErrorResponse(
			ctx, domain.NewAppError(http.StatusBadRequest, err),
		)
	}

	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.ErrorResponse(ctx, err)
	}

	err = p.servicePayments.Create(ctx.Request().Context(), userID, paymentRequest)
	if err != nil {
		return domain.ErrorResponse(ctx, err)
	}

	return ctx.NoContent(http.StatusCreated)
}

func (p *Payments) Get(ctx *echo.Context) error {
	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return domain.ErrorResponse(
			ctx, domain.NewAppError(http.StatusBadRequest, err),
		)
	}
	payments, err := p.servicePayments.Get(ctx.Request().Context(), userID)
	if err != nil {
		return domain.ErrorResponse(ctx, err)
	}

	return ctx.JSON(http.StatusOK, payments)
}
