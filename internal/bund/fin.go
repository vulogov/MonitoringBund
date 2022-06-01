package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/vulogov/monitoringbund/internal/banner"
)

func Fin() {
	banner.Banner("[ Zay Gezunt ]")
	log.Debugf("[ MBUND ] bund.Fin(%v) is reached", ApplicationId)
	if NewRelicConfigured {
		NRAPI.Close()
	}
	CloseNatsAgent()
	CloseEtcdAgent()
	CloseInternalStorage()
	log.Infof("[ MBUND ] %s is down", ApplicationId)
	signal.ExitRequest()
}
