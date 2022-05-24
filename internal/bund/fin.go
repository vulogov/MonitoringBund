package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/vulogov/monitoringbund/internal/banner"
)

func Fin() {
	banner.Banner("[ Zay Gezunt ]")
	log.Debugf("[ MBUND ] bund.Fin(%v) is reached", ApplicationId)
	CloseNatsAgent()
	CloseEtcdAgent()
	log.Debug("Wait while NR application is shut down")
	log.Infof("[ MBUND ] %s is down", ApplicationId)
	signal.ExitRequest()
}
