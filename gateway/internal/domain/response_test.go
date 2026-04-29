package domain_test

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/SynKolbasyn/bank/gateway/internal/model"
	"github.com/SynKolbasyn/bank/gateway/pkg/logger"
	"github.com/labstack/echo/v5"
	"github.com/stretchr/testify/require"
)

func TestErrorResponse(t *testing.T) {
	t.Parallel()

	testData := []int{
		http.StatusOK,
		http.StatusBadRequest,
		http.StatusUnauthorized,
		http.StatusForbidden,
		http.StatusNotFound,
		http.StatusConflict,
		http.StatusUnprocessableEntity,
		http.StatusInternalServerError,
		http.StatusNotImplemented,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
	}

	server := echo.New()
	server.Logger = logger.NewLogger(slog.LevelError + 2)

	for i, data := range testData {
		t.Run(fmt.Sprintf("error-%d", i+1), func(t *testing.T) {
			t.Parallel()

			require.NotNil(t, data)

			req, err := http.NewRequestWithContext(t.Context(), "", "", nil)
			require.NoError(t, err)
			require.NotNil(t, req)

			rec := httptest.NewRecorder()
			require.NotNil(t, rec)

			ctx := server.NewContext(req, rec)

			require.NoError(t, domain.ErrorResponse(ctx, domain.NewAppError(data)))

			require.Equal(t, data, rec.Code)

			var resp model.ErrorResponse
			require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
			require.Equal(t, resp.Error, http.StatusText(data))
		})
	}
}

func TestUnknownErrorResponse(t *testing.T) {
	t.Parallel()

	server := echo.New()
	server.Logger = logger.NewLogger(slog.LevelError + 2)

	testCode := http.StatusInternalServerError

	testData := []error{ErrTest, nil}

	for i, data := range testData {
		t.Run(fmt.Sprintf("ErrorResponse-%d", i+1), func(t *testing.T) {
			t.Parallel()
			req, err := http.NewRequestWithContext(t.Context(), "", "", nil)
			require.NoError(t, err)
			require.NotNil(t, req)

			rec := httptest.NewRecorder()
			require.NotNil(t, rec)

			ctx := server.NewContext(req, rec)

			require.NoError(t, domain.ErrorResponse(ctx, data))

			require.Equal(t, testCode, rec.Code)

			var resp model.ErrorResponse
			require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
			require.Equal(t, resp.Error, http.StatusText(testCode))
		})
	}
}
