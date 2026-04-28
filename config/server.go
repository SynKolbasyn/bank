package config

import (
	"log/slog"
	"net"
)

type Server struct {
	host string
	port string
	LogLevel slog.Level
}

func LoadServer() (*Server, error) {
	var logLevel slog.Level
	err := logLevel.UnmarshalText([]byte(KeyServerLogLevel.GetValueDefault("INFO")))
	if err != nil {
		return nil, err
	}
	server := &Server{
		host: KeyServerHost.GetValueDefault("0.0.0.0"),
		port: KeyServerPort.GetValueDefault("80"),
		LogLevel: logLevel,
	}
	return server, nil
}

func (s *Server) Address() string {
	return net.JoinHostPort(s.host, s.port)
}
