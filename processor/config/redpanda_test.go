package config_test

import (
	"fmt"
	"testing"

	"github.com/SynKolbasyn/bank/processor/config"
	"github.com/stretchr/testify/require"
)

func TestLoadRedpanda(t *testing.T) {
	testData := []struct{
		Hosts string
		Expected []string
	} {
		{"localhost", []string{"localhost"}},
		{"localhost:9092", []string{"localhost:9092"}},
		{"first,second", []string{"first", "second"}},
		{"first:9092,second:9092", []string{"first:9092", "second:9092"}},
		{"first:9092,second:9092,third", []string{"first:9092", "second:9092", "third"}},
	}

	for i, data := range testData {
		t.Run(fmt.Sprintf("hosts-%d", i + 1), func(t *testing.T) {
			t.Setenv(string(config.KeyRedpandaHosts), data.Hosts)

			redpanda := config.LoadRedpanda()
			require.NotNil(t, redpanda)
			require.Equal(t, data.Expected, redpanda.Hosts)
		})
	}
}
