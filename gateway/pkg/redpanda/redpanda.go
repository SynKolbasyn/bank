package redpanda

import "github.com/twmb/franz-go/pkg/kgo"

func NewClient(hosts []string, topics []string) (*kgo.Client, error) {
	opts := []kgo.Opt{kgo.SeedBrokers(hosts...)}
	if len(topics) != 0 {
		opts = append(opts, kgo.ConsumeTopics(topics...))
	}
	return kgo.NewClient(opts...)
}
