package app

import (
	"context"
	"errors"
	"log/slog"
	"runtime"

	"github.com/SynKolbasyn/bank/processor/internal/handler"
	"github.com/twmb/franz-go/pkg/kgo"
	"golang.org/x/sync/errgroup"
)


func StartProcessing(ctx context.Context, client *kgo.Client, handler *handler.Payments) {
	group, ctx := errgroup.WithContext(ctx)
	group.SetLimit(runtime.NumCPU())

	for {
		fetches := client.PollFetches(ctx)
		var err error
		for _, fetchErr := range fetches.Errors() {
			err = errors.Join(err, fetchErr.Err)
		}
		if err != nil {
			if errors.Is(err, context.Canceled) {
				break
			} else {
				slog.ErrorContext(ctx, "client.PollFetches", slog.String("error", err.Error()))
				continue
			}
		}
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			group.Go(func() error {
				err := handler.Process(ctx, record.Value)
				if err != nil {
					slog.ErrorContext(ctx, "handler.Process", slog.String("error", err.Error()))
				} else {
					err = client.CommitRecords(ctx, record)
					if err != nil {
						slog.ErrorContext(ctx, "client.CommitRecords", slog.String("error", err.Error()))
					}
				}
				return nil
			})
		}
	}
	group.Wait()
}
