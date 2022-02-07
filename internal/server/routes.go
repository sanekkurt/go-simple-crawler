package server

import (
	"github.com/go-chi/chi"
	"go-simple-crawler/internal/registries/server"
	"go-simple-crawler/internal/server/common"
)

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	common.ConfigHandlers(r)

	for _, v := range server.InitRegistry(s.log, s.cfg) {
		r.Route("/"+v.GetName(), v.GetRoutes())
	}

	return r
}
