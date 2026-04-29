package config_test

import (
	"testing"

	"github.com/SynKolbasyn/bank/processor/config"
	"github.com/stretchr/testify/require"
)

func TestLoadAuth(t *testing.T) {
	auth := config.LoadAuth()
	require.NotNil(t, auth)
	require.Equal(t, auth.Secret, []byte("auth-secret"))

	t.Setenv("AUTH_SECRET", "very-good-auth-secret")

	auth = config.LoadAuth()
	require.NotNil(t, auth)
	require.Equal(t, auth.Secret, []byte("very-good-auth-secret"))
}
