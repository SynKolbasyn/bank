package middleware

import (
	"net/http"

	"github.com/SynKolbasyn/bank/internal/domain"
	"github.com/SynKolbasyn/bank/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

const KeyToken string = "user"

func NewJWTConfig(secretKey []byte) echojwt.Config {
	//nolint:gosec
	return echojwt.Config{
		ErrorHandler: func(ctx *echo.Context, err error) error {
			return domain.ErrorResponse(ctx, domain.NewAppError(http.StatusUnauthorized, err))
		},
		SigningKey:  secretKey,
		ContextKey:  KeyToken,
		NewClaimsFunc: func(_ *echo.Context) jwt.Claims {
			return new(model.JWTAuthData)
		},
	}
}

func GetJWTData(ctx *echo.Context) (*model.JWTAuthData, error) {
	token, err := echo.ContextGet[*jwt.Token](ctx, KeyToken)
	if err != nil {
		return nil, domain.NewAppError(http.StatusUnauthorized, err)
	}

	jwtData, ok := token.Claims.(*model.JWTAuthData)
	if !ok {
		return nil, domain.NewAppError(http.StatusForbidden, nil)
	}

	return jwtData, nil
}

func GetUserID(ctx *echo.Context) (uuid.UUID, error) {
	jwtData, err := GetJWTData(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}

	return jwtData.UserID, nil
}
