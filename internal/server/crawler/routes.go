package crawler

import (
	"github.com/go-chi/chi"
	"go-simple-crawler/internal/config"
	"go.uber.org/zap"
)

func Init(l *zap.SugaredLogger, cfg config.AppConfig) Service {
	return Service{log: l, cfg: cfg}
}

type Service struct {
	log *zap.SugaredLogger
	cfg config.AppConfig
}

func (s Service) GetRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", s.post)
	}
}

func (s Service) GetName() string {
	return "crawler"
}
