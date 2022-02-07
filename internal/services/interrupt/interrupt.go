package interrupt

import (
	"go.uber.org/zap"
	"os"
)

func WaitInterruptSignal(log *zap.SugaredLogger, c chan os.Signal) {
	for val := range c {
		log.Infof("Got signal: %s. Graceful shutdown", val)
	}
}
