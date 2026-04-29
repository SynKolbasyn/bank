package domain_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/stretchr/testify/require"
)

var ErrTest = errors.New("test error")

func TestAppError(t *testing.T) {
	t.Parallel()

	err := domain.NewAppError(http.StatusInternalServerError, ErrTest)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrTest)
	require.Equal(t, ErrTest.Error(), err.Error())
}

func TestNilAppError(t *testing.T) {
	t.Parallel()

	err := domain.NewAppError(http.StatusInternalServerError)
	require.Error(t, err)
	require.Empty(t, err.Error())
}
