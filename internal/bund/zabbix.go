package bund

import (
	"fmt"
	"math"
	"time"
	"os"
	"strconv"
	zbxagent "github.com/akomic/go-zabbix-proto/agent"
	"github.com/lestrrat-go/strftime"
	"github.com/Jeffail/gabs/v2"
	"github.com/nats-io/nats.go"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
)

var ZabbixConfigured bool


func InitZabbixAgent() bool {
	if len(*conf.ZBX_account) == 0 {
		if acc, ok := Conf.Load("ZABBIX_ACCOUNT"); ok {
			*conf.ZBX_account = string(acc.([]uint8))
		}
	}
	if len(*conf.ZBX_pass) == 0 {
		if zpass, ok := Conf.Load("ZABBIX_PASSWORD"); ok {
			*conf.ZBX_pass = string(zpass.([]uint8))
		}
	}
	if len(*conf.ZBX_api) == 0 {
		if api, ok := Conf.Load("ZABBIX_API"); ok {
			*conf.ZBX_api = string(api.([]uint8))
		}
	}
	if len(*conf.ZBX_host) == 0 {
		if host, ok := Conf.Load("ZABBIX_HOST"); ok {
			*conf.ZBX_host = string(host.([]uint8))
		}
	}
	if len(*conf.ZBX_port) == 0 {
		if port, ok := Conf.Load("ZABBIX_PORT"); ok {
			*conf.ZBX_port = string(port.([]uint8))
		}
	}
	_, err := strconv.Atoi(*conf.ZBX_port)
	if err != nil {
		log.Errorf("[ ZABBIX ] Port value is incorrect: %v", *conf.ZBX_port)
		return false
	}
	if len(*conf.ZBX_account) > 0 && len(*conf.ZBX_pass) > 0 && len(*conf.ZBX_api) > 0 && len(*conf.ZBX_host) > 0 && len(*conf.ZBX_port) > 0 {
		ZabbixConfigured = true
	}
	if ! ZabbixConfigured {
		log.Error("ZABBIX environment is not configured")
	} else {
		log.Debug("ZABBIX environment is configured")
	}
	return ZabbixConfigured
}

func SendZabbixMetric(data *gabs.Container) {
	var zdata []*zbxagent.Metric
	var res *zbxagent.Response
	var err error

	host  := data.Search("metric", "host").Data().(string)
	key   := data.Search("metric", "key").Data().(string)
	value := data.Search("metric", "value").Data()
	stamp := int64(data.Search("metric", "timestamp").Data().(float64) / 1000)
	if value == nil {
		value = math.NaN()
	}
	zbx_port, _ := strconv.Atoi(*conf.ZBX_port)
	agent := zbxagent.NewAgent(host, *conf.ZBX_host, zbx_port)
	switch zvalue := value.(type) {
	case string, int64, int32, float32, float64, int:
		zdata = append(zdata, agent.NewMetric(key, fmt.Sprintf("%v", zvalue), stamp))
	default:
		return
	}

  res, err = agent.Send(zdata)
	if err != nil {
    log.Errorf("[ ZABBIX ] Error sending data: %v", err)
  } else {
    log.Debugf("[ ZABBIX ] OK: %v", res.JSON)
  }
}

func SendZabbixEvent(data *gabs.Container) {
	var zdata []*zbxagent.Metric
	var res *zbxagent.Response
	var err error

	host  := data.Search("event", "host").Data().(string)
	key   := data.Search("event", "key").Data().(string)
	value := data.Search("attributes", "value").Data()
	if value == nil {
		return
	}
	zbx_port, _ := strconv.Atoi(*conf.ZBX_port)
	agent := zbxagent.NewAgent(host, *conf.ZBX_host, zbx_port)
	switch zvalue := value.(type) {
	case string, int64, int32, float32, float64, int:
		stamp := int64(data.Search("event", "timestamp").Data().(float64) / 1000)
		zdata = append(zdata, agent.NewMetric(key, fmt.Sprintf("%v", zvalue), stamp))
	default:
		return
	}

  res, err = agent.Send(zdata)
	if err != nil {
    log.Errorf("[ ZABBIX ] Error sending data: %v", err)
  } else {
    log.Debugf("[ ZABBIX ] OK: %v", res.JSON)
  }
}

func SendZabbixLog(data *gabs.Container) {
	var zdata []*zbxagent.Metric
	var res *zbxagent.Response
	var err error
	var datebuf string

	host  := data.Search("log", "host").Data().(string)
	svc   := data.Search("log", "service").Data().(string)
	lt   	:= data.Search("log", "logtype").Data().(string)
	msg 	:= data.Search("log", "msg").Data()

	zbx_port, _ := strconv.Atoi(*conf.ZBX_port)
	agent := zbxagent.NewAgent(host, *conf.ZBX_host, zbx_port)
	switch zvalue := msg.(type) {
	case string:
		key   := fmt.Sprintf("%v.%v", svc, lt)
		stamp := int64(data.Search("log", "timestamp").Data().(float64) / 1000)
		f1, err := strftime.New("%Y%m%d:%H%M%S.%L", strftime.WithMilliseconds('L'))
		if datebuf = f1.FormatString(time.Now()); err != nil {
			log.Errorf("[ ZABBIX ] date format error: %v", err)
			return
		}
		zmsg  := fmt.Sprintf("%05d:%v %v", os.Getpid(), datebuf, zvalue)
		zdata = append(zdata, agent.NewMetric(key, zmsg, stamp))
	default:
		return
	}

	res, err = agent.Send(zdata)
	if err != nil {
		log.Errorf("[ ZABBIX ] Error sending data: %v", err)
	} else {
		log.Debugf("[ ZABBIX ] OK: %v", res.JSON)
	}

}

func ProcessZabbixMetric(m *nats.Msg) {
	if ! ZabbixConfigured{
		log.Debug("Packet received, but Zabbix not configured")
		return
	}
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	data, err := gabs.ParseJSON(msg.Value)
	if err != nil {
		log.Errorf("[ ZABBIX ] Packet parsing error: %v", err)
		return
	}
	t := data.Search("type").Data()
	if t == nil {
		log.Errorf("[ ZABBIX ] Malformatted packet: %v", data.String())
		return
	}
	if msg.PktClass == "TELEMETRY" {
		switch msg.PktKey {
		case "METRIC":
			if t.(string) == "metric" {
				SendZabbixMetric(data)
			}
		case "EVENT":
			if t.(string) == "event" {
				SendZabbixEvent(data)
			}
		case "LOG":
			if t.(string) == "log" {
				SendZabbixLog(data)
			}
		default:
			log.Debugf("[ ZABBIX ] Invalid packet %v[%v,%v]", msg.PktId, msg.PktClass, msg.PktKey)
		}
	}
}

func Zabbix_Client() {
	Init()
	InitEtcdAgent("zabbix_client")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! InitZabbixAgent() {
		return
	}
	if ! WaitSync() {
		return
	}
	log.Debugf("[ MBUND ] bund.Zabbix_Client(%v) is reached", ApplicationId)
	NatsTelemetryRecv(ProcessZabbixMetric)
	Loop()
}

func init() {
	ZabbixConfigured = false
}
