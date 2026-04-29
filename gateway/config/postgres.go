package config

import (
	"net"
	"net/url"
)

type Postgres struct {
	host     string
	port     string
	user     string
	password string
	database string
}

func LoadPostgres() *Postgres {
	return &Postgres{
		host:     KeyPostgresHost.GetValueDefault("localhost"),
		port:     KeyPostgresPort.GetValueDefault("5432"),
		user:     KeyPostgresUser.GetValueDefault("user"),
		password: KeyPostgresPassword.GetValueDefault("password"),
		database: KeyPostgresDatabase.GetValueDefault("database"),
	}
}

func (p *Postgres) DSN() string {
	url := url.URL{
		Scheme: "postgres",
		User: url.UserPassword(p.user, p.password),
		Host: net.JoinHostPort(p.host, p.port),
		Path: p.database,
	}
	return url.String()
}
