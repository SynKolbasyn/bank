package config

type Config struct {
	Server *Server
	Auth *Auth
	Postgres *Postgres
	Redpanda *Redpanda
}

func LoadConfig() (*Config, error) {
	server, err := LoadServer()
	if err != nil {
		return nil, err
	}
	auth := LoadAuth()
	postgres := LoadPostgres()
	redpanda := LoadRedpanda()

	config := &Config{
		Server: server,
		Auth: auth,
		Postgres: postgres,
		Redpanda: redpanda,
	}
	return config, nil
}
