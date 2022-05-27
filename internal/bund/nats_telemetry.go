package bund

import (
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
)



func NatsPublishTelemetry(QueueName string, data []byte) {
	if DoContinue && HadSync {
		log.Debugf("[ PUB ] %v", QueueName)
		Nats.Publish(QueueName, data)
	} else {
		log.Errorf("[ PUB ] not ready %v", QueueName)
	}
}

func NatsSendTelemetry(data *tc.TCJson, uri string, pktkey string) error {
	jdata := data.J.Bytes()
	pkt, err := Marshal(uri, ApplicationType, "TELEMETRY", pktkey, nil, jdata, nil)
	if err != nil {
		return err
	}
	NatsPublishTelemetry(uri, pkt)
	return nil
}

func NatsSendEvent(data *tc.TCJson) error {
	return NatsSendTelemetry(data, EvtQueueName, "EVENT")
}

func NatsSendMetric(data *tc.TCJson) error {
	return NatsSendTelemetry(data, MetricQueueName, "METRIC")
}

func NatsSendLog(data *tc.TCJson) error {
	return NatsSendTelemetry(data, LogQueueName, "LOG")
}

func NatsSendTrace(data *tc.TCJson) error {
	return NatsSendTelemetry(data, TraceQueueName, "TRACE")
}
