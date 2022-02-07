package app

import (
	"context"
	"fmt"
	"go-simple-crawler/internal/services/interrupt"
	"os"
	"os/signal"

	"go-simple-crawler/internal/config"
	"go-simple-crawler/internal/logging"
	"go-simple-crawler/internal/server"
)

func RunApp() {
	var (
		ctx = context.Background()
	)

	log, err := logging.Configure(false)
	if err != nil {
		fmt.Println("[ERROR] ", err.Error())
		os.Exit(2)
	}

	ctx = logging.WithLogger(ctx, log)

	cfg, err := config.Parse(ctx, os.Args)
	if err != nil {
		if err == config.ErrHelpShown {
			return
		}

		log.Error(err)
		return
	}

	log.Info("go-simple-crawler server starting")

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go interrupt.WaitInterruptSignal(log, c, cancelFunc)

	srv, err := server.NewServer(ctx, cfg, log)
	if err != nil {
		log.Errorf("failed to get the server: %s", err)
		return
	}

	srv.Run(ctx)
}
