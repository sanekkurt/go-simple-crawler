package interrupt

import (
	"context"
	"go.uber.org/zap"
	"os"
)

func WaitInterruptSignal(log *zap.SugaredLogger, c chan os.Signal, cancel context.CancelFunc) {
	val := <-c

	log.Infof("Got signal: %s. Graceful shutdown", val)
	cancel()
}
