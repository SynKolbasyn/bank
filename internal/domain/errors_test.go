package domain_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/SynKolbasyn/bank/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestAppError(t *testing.T) {
	t.Parallel()

	testErr := errors.New("test error")

	err := domain.NewAppError(http.StatusInternalServerError, testErr)
	require.Error(t, err)
	require.Equal(t, testErr.Error(), err.Error())
	require.ErrorIs(t, err, testErr)
}

func TestNilAppError(t *testing.T) {
	t.Parallel()

	err := domain.NewAppError(http.StatusInternalServerError)
	require.Error(t, err)
	require.Empty(t, err.Error())
}
