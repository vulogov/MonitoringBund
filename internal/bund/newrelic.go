package bund

import (
	"fmt"
	"math"
	"strings"
	"github.com/Jeffail/gabs/v2"
	"github.com/vulogov/nrapi"
	"github.com/nats-io/nats.go"
	"github.com/pieterclaerhout/go-log"
	"github.com/peterh/liner"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/stdlib"
)

var NewRelicConfigured bool
var NRAPI *nrapi.NRAPI

var (
	nrql_commands = []string{
		".version", ".exit", ".stack", ".last", "SELECT", "FROM",
	}
	NRQL_PROMPT = "[ NRQL ] "
)

func InitNewRelicAgent() bool {
	if len(*conf.NR_account) == 0 {
		if acc, ok := Conf.Load("NEWRELIC_ACCOUNT"); ok {
			*conf.NR_account = string(acc.([]uint8))
		}
	}
	if len(*conf.NR_lic_key) == 0 {
		if lic, ok := Conf.Load("NEWRELIC_LICENSE_KEY"); ok {
			*conf.NR_lic_key = string(lic.([]uint8))
		}
	}
	if len(*conf.NR_api_key) == 0 {
		if api, ok := Conf.Load("NEWRELIC_API_KEY"); ok {
			*conf.NR_api_key = string(api.([]uint8))
		}
	}
	if len(*conf.NR_account) > 0 && len(*conf.NR_lic_key) > 0 && len(*conf.NR_api_key) > 0 {
		NewRelicConfigured = true
	}
	if ! NewRelicConfigured {
		log.Error("NEW RELIC environment is not configured")
	} else {
		log.Debug("NEW RELIC environment is configured")
		NRAPI = nrapi.New(*conf.NR_account, *conf.NR_lic_key, *conf.NR_api_key)
		if *conf.Debug {
			NRAPI.SetConsoleLog()
		} else {
			NRAPI.DisableLog()
		}
	}
	return NewRelicConfigured
}

func SendNewRelicMetric(data *gabs.Container) {
	attrs := TelemetryAttributesToMap(data)
	host  := data.Search("metric", "host").Data().(string)
	key   := data.Search("metric", "key").Data().(string)
	mtype := data.Search("metric", "type").Data().(string)
	value := data.Search("metric", "value").Data()
	if value == nil {
		value = math.NaN()
	}
	NRAPI.Metric(host, key, value, mtype, attrs)
}

func SendNewRelicEvent(data *gabs.Container) {
	attrs := TelemetryAttributesToMap(data)
	host  := data.Search("event", "host").Data().(string)
	key   := data.Search("event", "key").Data().(string)
	dst   := data.Search("event", "destination").Data().(string)
	value := data.Search("attributes", "value").Data()
	if value == nil {
		value = math.NaN()
	}
	NRAPI.Event(host, dst, key, value, attrs)
}


func ProcessMetric(m *nats.Msg) {
	if ! NewRelicConfigured {
		log.Debug("Packet received, but NewRelic API not configured")
		return
	}
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	data, err := gabs.ParseJSON(msg.Value)
	if err != nil {
		log.Errorf("[ NEWRELIC ] Packet parsing error: %v", err)
		return
	}
	t := data.Search("type").Data()
	if t == nil {
		log.Errorf("[ NEWRELIC ] Malformatted packet: %v", data.String())
		return
	}
	if msg.PktClass == "TELEMETRY" {
		switch msg.PktKey {
		case "METRIC":
			if t.(string) == "metric" {
				SendNewRelicMetric(data)
			}
		case "EVENT":
			if t.(string) == "event" {
				SendNewRelicEvent(data)
			}
		default:
			log.Debugf("[ NEWRELIC ] Invalid packet %v[%v,%v]", msg.PktId, msg.PktClass, msg.PktKey)
		}
	}
}

func Newrelic_Client() {
	Init()
	InitEtcdAgent("newrelic_client")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! InitNewRelicAgent() {
		return
	}
	if ! WaitSync() {
		return
	}
	log.Debugf("[ MBUND ] bund.Newrelic_Client(%v) is reached", ApplicationId)
	NatsTelemetryRecv(ProcessMetric)
	Loop()
	if NewRelicConfigured {
		NRAPI.Close()
	}
}

func Newrelic_NRQL_Shell() {
	Init()
	InitEtcdAgent("nrql")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! InitNewRelicAgent() {
		return
	}
	log.Debugf("[ MBUND ] bund.Newrelic_NRQL_Shell(%v) is reached", ApplicationId)
	log.Info("For exit, type: .exit")
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range nrql_commands {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	core := stdlib.InitBUND()

	out:
	for {
		if cmd, err := line.Prompt(NRQL_PROMPT); err == nil {
			cmd = strings.Trim(cmd, "\n \t\r")
			line.AppendHistory(cmd)
			log.Debugf("shell get: %v", cmd)
			switch cmd {
			case ".exit":
				log.Debug("Exiting")
				break out
			default:
				if IsShellCommand(cmd) {
					log.Debugf("Running shell command: %v", cmd)
					RunShellCommand(cmd, core.TC)
				} else {
					log.Debug("Sending NRQL")
					df, err := nrapi.DataFrame(NRAPI.NRQL(cmd))
					if err != nil {
						log.Errorf("[ NRQL ] %v", err)
						continue
					}
					fmt.Println(df.String())
					if core.TC.ExitRequested() {
						log.Debug("Exiting from shell")
						break out
					}
				}
			}
		} else if err == liner.ErrPromptAborted {
			log.Debug("Aborted")
			break
		} else {
			log.Debugf("Error reading line: %v", err)
		}
	}
	if NewRelicConfigured {
		NRAPI.Close()
	}
}

func init() {
	NewRelicConfigured = false
}
