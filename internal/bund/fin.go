package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/vulogov/monitoringbund/internal/banner"
)

func Fin() {
	signal.ExitRequest()
	banner.Banner("[ Zay Gezunt ]")
	log.Debugf("[ MBUND ] bund.Fin(%v) is reached", ApplicationId)
	if NewRelicConfigured {
		NRAPI.Close()
	}
	CloseNatsAgent()
	CloseEtcdAgent()
	CloseInternalStorage()
	log.Debug("Waiting for application to quit")
	signal.Loop()
	log.Infof("[ MBUND ] %s is down", ApplicationId)
}
