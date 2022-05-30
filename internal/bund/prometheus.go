package bund

import (
	"fmt"
	"math"
	"strconv"
	"github.com/Jeffail/gabs/v2"
	"github.com/nats-io/nats.go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
)

var PrometheusConfigured bool


func InitPrometheusAgent() bool {
	if len(*conf.PR_url) == 0 {
		if acc, ok := Conf.Load("PROMETHEUS_PUSHER_URL"); ok {
			*conf.PR_url = string(acc.([]uint8))
		}
	}

	if len(*conf.PR_url) > 0  {
		PrometheusConfigured = true
	}
	if ! PrometheusConfigured {
		log.Error("PROMETHEUS environment is not configured")
	} else {
		log.Debug("PROMETHEUS environment is configured")
	}
	return PrometheusConfigured
}

func SendPrometheusMetric(data *gabs.Container) {
	var desc string

	host  := data.Search("metric", "host").Data().(string)
	key   := data.Search("metric", "key").Data().(string)
	value := data.Search("metric", "value").Data()
	if value == nil {
		value = math.NaN()
	}

	if data.Exists("attributes", "description") {
		desc = data.Search("attributes", "description").Data().(string)
	}

	prom_key := fmt.Sprintf("%v_%v_%v", *conf.Name, host, key)

	metric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: prom_key,
		Help: desc,
	})

	switch zvalue := value.(type) {
	case int64:
		metric.Set(float64(zvalue))
	case int:
		metric.Set(float64(zvalue))
	case int32:
		metric.Set(float64(zvalue))
	case float32:
		metric.Set(float64(zvalue))
	case float64:
		metric.Set(zvalue)
	case string:
		_val, err := strconv.ParseFloat(zvalue, 64)
		if err != nil {
			log.Errorf("[ PROMETHEUS ] Invalid value for a metric: %v", err)
			return
		}
		metric.Set(_val)
	default:
		return
	}
	if err := push.New(*conf.PR_url, *conf.Name).
		Collector(metric).
		Grouping("key", key).
		Grouping("host", host).
		Grouping("applicationid", ApplicationId).
		Push(); err != nil {
		log.Errorf("[ PROMETHEUS ] Error pushing metric: %v", err)
	}
}



func ProcessPrometheusMetric(m *nats.Msg) {
	if ! PrometheusConfigured {
		log.Debug("Packet received, but Prometheus not configured")
		return
	}
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	data, err := gabs.ParseJSON(msg.Value)
	if err != nil {
		log.Errorf("[ PROMETHEUS ] Packet parsing error: %v", err)
		return
	}
	t := data.Search("type").Data()
	if t == nil {
		log.Errorf("[ PROMETHEUS ] Malformatted packet: %v", data.String())
		return
	}
	if msg.PktClass == "TELEMETRY" {
		switch msg.PktKey {
		case "METRIC":
			if t.(string) == "metric" {
				SendPrometheusMetric(data)
			}
		default:
			log.Debugf("[ PROMETHEUS ] Invalid packet %v[%v,%v]", msg.PktId, msg.PktClass, msg.PktKey)
		}
	}
}

func Prometheus_Client() {
	Init()
	InitEtcdAgent("prometheus_client")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! InitPrometheusAgent() {
		return
	}
	if ! WaitSync() {
		return
	}
	log.Debugf("[ MBUND ] bund.Prometheus_Client(%v) is reached", ApplicationId)
	NatsTelemetryRecv(ProcessPrometheusMetric)
	Loop()
}

func init() {
	PrometheusConfigured = false
}
