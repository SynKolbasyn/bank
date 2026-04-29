package main

import (
	"context"
	"errors"
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

var ErrStatusCode = errors.New("resp.StatusCode != http.StatusOK")

func main() {
	err := check()
	if err != nil {
		os.Exit(1)
	}
}

func check() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.LoadServer()
	if err != nil {
		slog.ErrorContext(ctx, "config.LoadServer", slog.String("error", err.Error()))

		return err
	}

	slog.SetDefault(logger.NewLogger(cfg.LogLevel))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getURL(cfg.Address()), nil)
	if err != nil {
		slog.ErrorContext(ctx, "http.NewRequestWithContext", slog.String("error", err.Error()))

		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.ErrorContext(ctx, "http.DefaultClient.Do", slog.String("error", err.Error()))

		return err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			slog.WarnContext(ctx, "body.Close", slog.String("error", err.Error()))
		}
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		var data slog.Attr
		if err != nil {
			data = slog.String("error", err.Error())
		} else {
			data = slog.String("body", string(body))
		}

		slog.ErrorContext(
			ctx,
			ErrStatusCode.Error(),
			slog.Int("status_code", resp.StatusCode),
			slog.String("status", resp.Status),
			data,
		)

		return ErrStatusCode
	}

	return nil
}

func getURL(addr string) string {
	healthURL := url.URL{
		Scheme: "http",
		Host:   addr,
		Path:   "/health",
	}

	return healthURL.String()
}
