package bund

import (
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/davecgh/go-spew/spew"
	"github.com/nats-io/nats.go"
	"github.com/pieterclaerhout/go-log"
)

func WatchDisplay(m *nats.Msg) {
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	log.Debugf("[ PACKET ] %v", msg.PktId)
	spew.Dump(msg)
}

func Watch() {
	Init()
	InitEtcdAgent("watch")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! WaitSync() {
		return
	}
	log.Debugf("[ MBUND ] bund.Watch(%v) is reached", ApplicationId)
	if ! *conf.WTele {
		NatsRecv(WatchDisplay)
	} else {
		NatsTelemetryRecv(WatchDisplay)
	}
	Loop()
}
