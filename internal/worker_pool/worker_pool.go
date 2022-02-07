package worker_pool

import (
	"context"
	"go-simple-crawler/internal/logging"
	"sync"
)

type Worker struct {
	Number int
}

func (w Worker) RunTasks(ctx context.Context, wg *sync.WaitGroup, c chan func()) {
	var (
		log = logging.GetLoggerFromContext(ctx)
	)

	defer wg.Done()

	for task := range c {
		log.Debugf("worker %d: task in progress...", w.Number)
		task()
	}

	log.Debugf("worker %d finish", w.Number)
}
