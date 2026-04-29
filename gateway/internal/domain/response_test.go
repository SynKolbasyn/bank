package domain_test

import (
	"encoding/json"
	"errors"
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

	e := echo.New()
	e.Logger = logger.NewLogger(slog.LevelError + 2)

	for i, data := range testData {
		t.Run(fmt.Sprintf("error-%d", i + 1), func(t *testing.T) {
			t.Parallel()

			require.NotNil(t, data)

			r, err := http.NewRequestWithContext(t.Context(), "", "", nil)
			require.NoError(t, err)
			require.NotNil(t, r)

			w := httptest.NewRecorder()
			require.NotNil(t, w)

			ctx := e.NewContext(r, w)

			require.NoError(t, domain.ErrorResponse(ctx, domain.NewAppError(data)))

			require.Equal(t, data, w.Code)
			var resp model.ErrorResponse
			require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
			require.Equal(t, resp.Error, http.StatusText(data))
		})
	}
}

func TestUnknownErrorResponse(t *testing.T) {
	t.Parallel()

	e := echo.New()
	e.Logger = logger.NewLogger(slog.LevelError + 2)
	
	testErr := errors.New("some unknown error")
	testCode := http.StatusInternalServerError

	r, err := http.NewRequestWithContext(t.Context(), "", "", nil)
	require.NoError(t, err)
	require.NotNil(t, r)

	w := httptest.NewRecorder()
	require.NotNil(t, w)

	ctx := e.NewContext(r, w)

	require.NoError(t, domain.ErrorResponse(ctx, testErr))

	require.Equal(t, testCode, w.Code)
	var resp model.ErrorResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Equal(t, resp.Error, http.StatusText(testCode))
}

func TestNilErrorResponse(t *testing.T) {
	t.Parallel()

	e := echo.New()
	e.Logger = logger.NewLogger(slog.LevelError + 2)
	
	testCode := http.StatusInternalServerError

	r, err := http.NewRequestWithContext(t.Context(), "", "", nil)
	require.NoError(t, err)
	require.NotNil(t, r)

	w := httptest.NewRecorder()
	require.NotNil(t, w)

	ctx := e.NewContext(r, w)

	require.NoError(t, domain.ErrorResponse(ctx, nil))

	require.Equal(t, testCode, w.Code)
	var resp model.ErrorResponse
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Equal(t, resp.Error, http.StatusText(testCode))
}
