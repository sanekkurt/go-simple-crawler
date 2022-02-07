package config

import (
	"context"
	"fmt"
	"go-simple-crawler/internal/config/errors"
	"os"
	"path/filepath"

	"go-simple-crawler/internal/config/configstructs"
	"go-simple-crawler/internal/logging"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server configstructs.Server `yaml:"server"`
	Worker configstructs.Worker `yaml:"worker"`
}

type CtxKeyT string

var (
	ctxKey CtxKeyT = "configCtxKey"
)

func ParseConfig(ctx context.Context, cfgPath string) (AppConfig, error) {
	log := logging.GetLoggerFromContext(ctx)

	configFullPath, err := filepath.Abs(cfgPath)
	if err != nil {
		return AppConfig{}, fmt.Errorf("cannot resolve absolute path to %s: %w", configFullPath, err)
	}

	f, err := os.Open(configFullPath)
	if err != nil {
		return AppConfig{}, fmt.Errorf("cannot open config '%s' for reading: %w", configFullPath, err)
	}

	defer func() {
		err = f.Close()
		if err != nil {
			log.Warnf("cannot close config '%s': %s", cfgPath, err)
		}
	}()

	dec := yaml.NewDecoder(f)
	dec.SetStrict(true)

	cfg := AppConfig{}

	err = dec.Decode(&cfg)
	if err != nil {
		return AppConfig{}, fmt.Errorf("cannot parse config from '%s': %w", cfgPath, err)
	}

	if cfg.Server.Listen == 0 {
		return AppConfig{}, fmt.Errorf("%w. the listening port is not specified or is specified incorrectly", errors.ErrInvalidConfig)
	}

	if cfg.Worker.Count == 0 {
		return AppConfig{}, fmt.Errorf("%w. the worker count is not specified or is specified incorrectly", errors.ErrInvalidConfig)
	}

	return cfg, nil
}

func WithConfig(ctx context.Context, config AppConfig) context.Context {
	return context.WithValue(ctx, ctxKey, config)
}

func GetConfigFromContext(ctx context.Context) AppConfig {
	if config, ok := ctx.Value(ctxKey).(AppConfig); ok {
		return config
	}

	return AppConfig{}
}
