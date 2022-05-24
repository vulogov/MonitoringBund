package bund

import (
	"os"
	"time"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var NRapp newrelic.Application

func InitNewRelicAgent() {
	NRapp, err := newrelic.NewApplication(
		newrelic.ConfigAppName(ApplicationId),
		newrelic.ConfigLicense(*conf.NRLicenseKey),
		newrelic.ConfigDebugLogger(os.Stdout),
	)
	if err != nil {
		log.Errorf("[ NEWRELIC ] %v", err)
		os.Exit(10)
	}
	if err := NRapp.WaitForConnection(15 * time.Second); nil != err {
		log.Errorf("[ NEWRELIC ] %v", err)
		os.Exit(10)
	}
	log.Debugf("NR application %v has been initialized and connection established", ApplicationId)
}
