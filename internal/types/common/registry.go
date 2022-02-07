package common

import "github.com/go-chi/chi"

type IServerServicesRegistry []IServerService

type IServerService interface {
	GetName() string
	GetRoutes() func(r chi.Router)
}
