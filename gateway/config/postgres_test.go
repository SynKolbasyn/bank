package config_test

import (
	"fmt"
	"testing"

	"github.com/SynKolbasyn/bank/gateway/config"
	"github.com/stretchr/testify/require"
)

func TestLoadPostgres(t *testing.T) {
	testData := []struct {
		Host     string
		Port     string
		User     string
		Password string
		Database string
		Expected string
	}{
		{
			"localhost",
			"5432",
			"user",
			"password",
			"database",
			"postgres://user:password@localhost:5432/database",
		},
		{
			"localhost",
			"5432",
			"user",
			"pass:word",
			"database",
			"postgres://user:pass%3Aword@localhost:5432/database",
		},
		{
			"localhost",
			"5432",
			"us:er",
			"password",
			"database",
			"postgres://us%3Aer:password@localhost:5432/database",
		},
		{
			"local:host",
			"5432",
			"user",
			"password",
			"database",
			"postgres://user:password@[local:host]:5432/database",
		},
	}

	for i, data := range testData {
		t.Run(fmt.Sprintf("dsn-%d", i+1), func(t *testing.T) {
			t.Setenv(string(config.KeyPostgresHost), data.Host)
			t.Setenv(string(config.KeyPostgresPort), data.Port)
			t.Setenv(string(config.KeyPostgresUser), data.User)
			t.Setenv(string(config.KeyPostgresPassword), data.Password)
			t.Setenv(string(config.KeyPostgresDatabase), data.Database)

			postgres := config.LoadPostgres()
			require.NotNil(t, postgres)
			require.Equal(t, data.Expected, postgres.DSN())
		})
	}
}
