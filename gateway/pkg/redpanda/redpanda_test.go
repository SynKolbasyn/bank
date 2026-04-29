package redpanda_test

import (
	"testing"

	"github.com/SynKolbasyn/bank/gateway/pkg/redpanda"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Parallel()

	brokers := []string{"redpanda"}
	topics := []string{"topic"}
	redpanda, err := redpanda.NewClient(brokers, topics)
	require.NoError(t, err)
	require.NotNil(t, redpanda)
	require.Len(t, redpanda.SeedBrokers(), len(brokers))
	require.Equal(t, topics, redpanda.GetConsumeTopics())
}

func TestNewClientNoTopics(t *testing.T) {
	t.Parallel()

	brokers := []string{"redpanda"}
	redpanda, err := redpanda.NewClient(brokers, nil)
	require.NoError(t, err)
	require.NotNil(t, redpanda)
	require.Len(t, redpanda.SeedBrokers(), len(brokers))
}

func TestNewClientNoHosts(t *testing.T) {
	t.Parallel()

	redpanda, err := redpanda.NewClient(nil, nil)
	require.Error(t, err)
	require.Nil(t, redpanda)
}
