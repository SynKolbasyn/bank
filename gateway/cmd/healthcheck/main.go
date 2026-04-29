package main

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/SynKolbasyn/bank/gateway/config"
	"github.com/SynKolbasyn/bank/gateway/pkg/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.LoadServer()
	if err != nil {
		slog.ErrorContext(ctx, "config.LoadServer", slog.String("error", err.Error()))
		os.Exit(1)
	}

	slog.SetDefault(logger.NewLogger(cfg.LogLevel))

	healthURL := url.URL{
		Scheme: "http",
		Host: cfg.Address(),
		Path: "/health",
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, healthURL.String(), nil)
	if err != nil {
		slog.ErrorContext(ctx, "http.NewRequestWithContext", slog.String("error", err.Error()))
		os.Exit(1)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "http.DefaultClient.Do", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer func (ctx context.Context, body io.ReadCloser)  {
		err := body.Close()
		if err != nil {
			slog.WarnContext(ctx, "body.Close", slog.String("error", err.Error()))
		}
	}(ctx, resp.Body)
    if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		var data slog.Attr
		if err != nil {
			data = slog.String("error", err.Error())
		} else {
			data = slog.String("body", string(body))
		}
		slog.ErrorContext(ctx, "resp.StatusCode != http.StatusOK", slog.Int("status_code", resp.StatusCode), slog.String("status", resp.Status), data)
		os.Exit(1)
	}
}
