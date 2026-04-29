package domain_test

import (
	"testing"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestValidator(t *testing.T) {
	t.Parallel()
	
	validator := domain.NewValidator()
	require.NotNil(t, validator)
	data := struct{
		password string `validate:"required,min=8,max=32"`
	}{
		"SuperPuper123",
	}
	err := validator.Validate(data)
	require.Nil(t, err)
}
