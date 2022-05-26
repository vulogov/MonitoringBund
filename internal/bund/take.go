package bund

import (
	"time"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/nats-io/nats.go"
	"github.com/pieterclaerhout/go-log"
)

func NRBundExecuteScript(m *nats.Msg) {
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	if msg.PktKey == "Agitator" && len(msg.Value) > 0 {
		log.Debugf("Script: %v", msg.Uri)
		BundGlobalEvalExpression(string(msg.Value), msg.Args, msg.Res)
	}
	if len(msg.Res) > 0 {
		log.Debug("Wait for NR events to catch up")
		time.Sleep(5 * time.Second)
	}
	signal.ExitRequest()
	DoContinue = false
}

func Take() {
	Init()
	InitEtcdAgent("take")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! WaitSync() {
		return
	}
	log.Debugf("[ MBUND ] bund.Take(%v) is reached", ApplicationId)
	NatsRecv(NRBundExecuteScript)
	Loop()
}
