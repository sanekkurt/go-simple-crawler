package config

import (
	"context"
	"errors"
	"fmt"

	"github.com/jessevdk/go-flags"
)

var (
	opts struct {
		ConfigPath string `long:"config" short:"c" env:"CONFIG_PATH" description:"Path to config.yaml file" required:"true"`
	}
	ErrHelpShown = errors.New("help message shown")
)

func Parse(ctx context.Context, args []string) (AppConfig, error) {
	_, err := flags.NewParser(&opts, flags.HelpFlag|flags.PassDoubleDash).ParseArgs(args[1:])
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok {
			if flagsErr.Type == flags.ErrHelp {
				return AppConfig{}, ErrHelpShown
			}

			return AppConfig{}, fmt.Errorf("cannot parse arguments: %w", flagsErr)
		}

		return AppConfig{}, fmt.Errorf("cannot parse arguments: %w", err)
	}

	cfg, err := ParseConfig(ctx, opts.ConfigPath)
	if err != nil {
		return AppConfig{}, err
	}

	return cfg, nil
}
