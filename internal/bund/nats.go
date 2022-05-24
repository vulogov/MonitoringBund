package bund

import (
	"os"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/pieterclaerhout/go-log"
)

var Nats *nats.Conn
var QueueName string
var SysQueueName string
var HadSync bool

func SysQueueHandler(m *nats.Msg) {
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	if IfSTOP(msg) {
		log.Infof("STOP(%v) message received", msg.PktId)
	}
	if IfSYNC(msg) {
		if ! HadSync {
			HadSync = true
			log.Infof("SYNC(%v) message triggered SYNC state for %v", msg.PktId, ApplicationId)
		}
	}
}

func InitNatsAgent() {
	var err error

	log.Debugf("Connecting to NATS")
	Nats, err = nats.Connect(
		*conf.Gnats,
		nats.Name(ApplicationId),
		nats.ReconnectWait(*conf.Timeout),
		nats.PingInterval(*conf.Timeout),
		nats.Timeout(*conf.Timeout),
	)
	if err != nil {
		log.Errorf("[ NATS ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	QueueName = fmt.Sprintf("%s:%s", *conf.Id, *conf.Name)
	SysQueueName = fmt.Sprintf("%s:%s:sys", *conf.Id, *conf.Name)
	log.Debugf("Queue name: %v", QueueName)
	log.Debugf("SysQueue name: %v", SysQueueName)
	NatsRecvSys(SysQueueHandler)
}

func NatsSend(data []byte) {
	if DoContinue {
		Nats.Publish(QueueName, data)
	}
}

func NatsSendSys(data []byte) {
	if DoContinue {
		Nats.Publish(SysQueueName, data)
	}
}

func NatsRecv(fun nats.MsgHandler) {
	Nats.QueueSubscribe(QueueName, *conf.Id, fun)
}

func NatsRecvSys(fun nats.MsgHandler) {
	Nats.Subscribe(SysQueueName, fun)
}

func CloseNatsAgent() {
	log.Debugf("Terminating and draining NATS session")
	Nats.Flush()
}

func init() {
	HadSync = false
}
