package crawler

import (
	"encoding/json"
	"go-simple-crawler/internal/server/observability"
	"go-simple-crawler/internal/server/util"
	crawler2 "go-simple-crawler/internal/services/crawler"
	"go-simple-crawler/internal/types/crawler"
	"net/http"
)

func (s Service) post(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	ctx = observability.WithContext(ctx, s.log, s.cfg)
	dec := json.NewDecoder(r.Body)

	input := struct {
		Data crawler.URLsData `json:"data"`
	}{}

	err := dec.Decode(&input)
	if err != nil {
		s.log.Errorf("fail decode input data %v", err)
		util.WriteError(w, err)

		return
	}

	result, err := crawler2.GetTitles(ctx, input.Data)
	if err != nil {
		s.log.Errorf("fail GetTitles %v", err)
		util.WriteError(w, err)

		return
	}

	util.WriteData(w, http.StatusAccepted, struct {
		Data []crawler.ResultData `json:"data"`
	}{result})
}
