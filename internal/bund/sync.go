package bund

import (
	"time"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/pieterclaerhout/go-log"
)

type NATSSync struct {
}

func (s NATSSync) Run() {
	SendSync()
}

func SendSync() {
	data, err := MakeSync(ApplicationType)
	if err != nil {
		log.Errorf("[ MBUND ] SYNC: %v", err)
	}
	NatsSendSys(data)
}

func WaitSync() bool {
	log.Debugf("Waiting for SYNC from cluster")
	if *conf.NoSync {
		HadSync = true
		log.Warn("Skipping SYNC with cluster")
		return true
	}
	c := 0
	for c < 30 {
		if HadSync {
			return true
		}
		if signal.ExitRequested() {
			return false
		}
		c += 1
		SendSync()
		time.Sleep(1*time.Second)
	}
	log.Error("Can not receive SYNC from cluster")
	return false
}

func Sync() {
	Init()
	InitEtcdAgent("sync")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	log.Debugf("[ MBUND ] bund.Sync(%v) is reached", ApplicationId)
	for ! signal.ExitRequested() {
		time.Sleep(1 * time.Second)
	}
}
