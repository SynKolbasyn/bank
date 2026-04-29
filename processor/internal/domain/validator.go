package domain

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v5"
)

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i any) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}
