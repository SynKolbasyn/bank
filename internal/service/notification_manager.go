package service

import (
	"context"
	"net/http"

	"github.com/SynKolbasyn/bank/internal/domain"
	"github.com/google/uuid"
	"github.com/twmb/franz-go/pkg/kgo"
)

type NotificationManager struct {
	clientRedpanda *kgo.Client
	topic string
}

func NewNotificationManager(clientRedpanda *kgo.Client, topic string) *NotificationManager {
	return &NotificationManager{
		clientRedpanda: clientRedpanda,
		topic: topic,
	}
}

func (n *NotificationManager) Notify(ctx context.Context, paymentID uuid.UUID) error {
	value := []byte(paymentID.String())

	record := &kgo.Record{
		Topic: n.topic,
		Value: value,
	}

	err := n.clientRedpanda.ProduceSync(ctx, record).FirstErr()
	if err != nil {
		return domain.NewAppError(http.StatusInternalServerError, err)
	}
	return nil
}
