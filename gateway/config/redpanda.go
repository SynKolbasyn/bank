package config

import "strings"

type Redpanda struct {
	Hosts []string
}

func LoadRedpanda() *Redpanda {
	hosts := KeyRedpandaHosts.GetValueDefault("localhost:9092")

	return &Redpanda{
		Hosts: strings.Split(hosts, ","),
	}
}
