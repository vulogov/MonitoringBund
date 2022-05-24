package bund

import (
	"time"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/pieterclaerhout/go-log"
)

func SendStop() {
	data, err := MakeStop("stop")
	if err != nil {
		log.Errorf("[ MBUND ] STOP: %v", err)
	}
	NatsSendSys(data)
	signal.ExitRequest()
}

func Stop() {
	Init()
	InitEtcdAgent("stop")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	log.Debugf("[ MBUND ] bund.Stop(%v) is reached", ApplicationId)
	SendStop()
	for ! signal.ExitRequested() {
		time.Sleep(1 * time.Second)
	}
}
