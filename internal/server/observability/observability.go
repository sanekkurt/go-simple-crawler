package observability

import (
	"context"
	"go-simple-crawler/internal/config"
	"go-simple-crawler/internal/logging"
	"go.uber.org/zap"
)

func WithContext(ctx context.Context, log *zap.SugaredLogger, cfg config.AppConfig) context.Context {
	ctx = logging.WithLogger(ctx, log)
	ctx = config.WithConfig(ctx, cfg)

	return ctx
}
