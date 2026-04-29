package config

import "github.com/SynKolbasyn/bank/processor/pkg/config"

const (
	KeyServerHost     = config.Key("SERVER_HOST")
	KeyServerPort     = config.Key("SERVER_PORT")
	KeyServerLogLevel = config.Key("SERVER_LOG_LEVEL")

	KeyAuthSecret = config.Key("AUTH_SECRET")

	KeyPostgresHost     = config.Key("POSTGRES_HOST")
	KeyPostgresPort     = config.Key("POSTGRES_PORT")
	KeyPostgresUser     = config.Key("POSTGRES_USER")
	KeyPostgresPassword = config.Key("POSTGRES_PASSWORD")
	KeyPostgresDatabase = config.Key("POSTGRES_DATABASE")

	KeyRedpandaHosts = config.Key("REDPANDA_HOSTS")
)
