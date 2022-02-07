package helpers

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"go-simple-crawler/internal/logging"
	"net/http"
)

func GetTagInfoFromWebPage(ctx context.Context, url string, tag string) (string, error) {
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

	result := doc.Find(tag).Text()

	return result, nil
}
