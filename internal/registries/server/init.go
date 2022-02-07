package server

import (
	"go-simple-crawler/internal/config"
	"go-simple-crawler/internal/server/crawler"
	"go-simple-crawler/internal/types/common"
	"go.uber.org/zap"
)

func InitRegistry(log *zap.SugaredLogger, cfg config.AppConfig) common.IServerServicesRegistry {
	var (
		reg = make(common.IServerServicesRegistry, 0)
		srv = crawler.Init(log, cfg)
	)

	reg = append(reg, srv)

	return reg
}
