package config

import "strings"

type Redpanda struct {
	Hosts  []string
	Topics []string
}

func LoadRedpanda() *Redpanda {
	hosts := KeyRedpandaHosts.GetValueDefault("localhost:9092")
	topics := KeyRedpandaHosts.GetValueDefault("payments")

	return &Redpanda{
		Hosts:  strings.Split(hosts, ","),
		Topics: strings.Split(topics, ","),
	}
}
