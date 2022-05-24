package bund

import (
	"os"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/signal"
)

func CheckNewRelic() {
	log.Debug("[ MBUND ] Checking if New Relic environment is configured.")
	if len(*conf.NRLicenseKey) != 0 {
		log.Debugf("[ NEWRELIC ] key: %s ", *conf.NRLicenseKey)
	} else {
		log.Error("[ NEWRELIC ] Environment is not configured.")
		signal.ExitRequest()
		os.Exit(1)
	}
}
