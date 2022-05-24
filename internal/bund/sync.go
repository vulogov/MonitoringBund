package bund

import (
	"time"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/bamzi/jobrunner"
	"github.com/pieterclaerhout/go-log"
)

type NATSSync struct {
}

func (s NATSSync) Run() {
	log.Debug("[ MBUND ] sending Sync")
	SendSync()
}

func SendSync() {
	data, err := MakeSync("sync")
	if err != nil {
		log.Errorf("[ MBUND ] SYNC: %v", err)
	}
	NatsSendSys(data)
}

func Sync() {
	Init()
	InitEtcdAgent("sync")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	InitNewRelicAgent()
	log.Debugf("[ MBUND ] bund.Sync(%v) is reached", ApplicationId)
	jobrunner.Schedule("@every 5s", NATSSync{})
	for ! signal.ExitRequested() {
		time.Sleep(1 * time.Second)
	}
}
