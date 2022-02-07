package crawler

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-simple-crawler/internal/config"
	"go-simple-crawler/internal/logging"
	"go-simple-crawler/internal/types/crawler"
	wp "go-simple-crawler/internal/worker_pool"
	"net/http"
	"sync"
)

func GetTitles(ctx context.Context, data crawler.URLsData) ([]crawler.ResultData, error) {
	var (
		cfg = config.GetConfigFromContext(ctx)

		tasksChan        = make(chan func(), cfg.Worker.Count)
		tasksResultsChan = make(chan crawler.ResultData, cfg.Worker.Count)

		wg sync.WaitGroup
	)

	for i := 1; i <= cfg.Worker.Count; i++ {
		worker := wp.Worker{Number: i}
		wg.Add(1)
		go worker.RunTasks(ctx, &wg, tasksChan)
	}

	go func() {
		for _, url := range data {
			tasksChan <- getFunc(ctx, url, tasksResultsChan)
		}

		close(tasksChan)

	}()

	resChan := make(chan []crawler.ResultData)

	go func(input chan crawler.ResultData, output chan []crawler.ResultData) {
		var (
			res []crawler.ResultData
		)

		for r := range input {
			res = append(res, r)
		}

		output <- res
	}(tasksResultsChan, resChan)

	wg.Wait()
	close(tasksResultsChan)

	return <-resChan, nil

}

func getFunc(ctx context.Context, url string, c chan crawler.ResultData) func() {
	return func() {
		title, err := GetTitle(ctx, url)
		if err != nil {
			c <- crawler.ResultData{
				URL:   url,
				Error: err.Error(),
			}

			return
		}

		c <- crawler.ResultData{
			URL:   url,
			Title: title,
		}
	}
}

func GetTitle(ctx context.Context, url string) (string, error) {
	var (
		log = logging.GetLoggerFromContext(ctx)
	)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Error(err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error the page gave the wrong status: %d", res.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := doc.Find("title").Text()

	return result, nil
}
