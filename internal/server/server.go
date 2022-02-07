package server

import (
	"context"
	"errors"
	"fmt"
	"go-simple-crawler/internal/config"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	shutdownTimeout = 5 * time.Second
)

type Server struct {
	httpServer *http.Server
	listen     int
	log        *zap.SugaredLogger
	cfg        config.AppConfig
}

func NewServer(ctx context.Context, cfg config.AppConfig, log *zap.SugaredLogger) (*Server, error) {
	var (
		srv = &Server{
			listen: cfg.Server.Listen,
			log:    log,
			cfg:    cfg,
		}
	)

	srv.httpServer = &http.Server{
		Addr:              fmt.Sprintf(":%d", srv.listen),
		Handler:           srv.routes(),
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	return srv, nil
}

func (s *Server) Run(ctx context.Context) {
	s.log.Infof("server started and listen on '%d' port", s.listen)

	go func() {
		<-ctx.Done()

		s.log.Infof("shutdown initiated")

		shutdownCtx, done := context.WithTimeout(context.Background(), shutdownTimeout)

		defer done()

		if err := s.httpServer.Shutdown(shutdownCtx); err != nil {
			s.log.Errorf("http shutdown error: %s", err)
			return
		}

		s.log.Debugf("shutdown completed")
	}()

	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.log.Errorf("failed to serve: %s", err)
		return
	}

	s.log.Info("stop http server")
}
