package bund

import (
	"fmt"
	"time"
	"github.com/nats-io/nats.go"
	"github.com/pieterclaerhout/go-log"
	"github.com/Jeffail/gabs/v2"
	"github.com/nakabonne/tstorage"
	"github.com/vulogov/monitoringbund/internal/signal"
)

const MAXPIPECAP = 1000000

var StoragePipe  chan *gabs.Container

func StoreTelemetry(m *nats.Msg) {
	msg := UnMarshal(m.Data)
	if msg == nil {
		log.Error("Invalid packet received")
	}
	data, err := gabs.ParseJSON(msg.Value)
	if err != nil {
		log.Errorf("[ CATCHER ] Packet parsing error: %v", err)
		return
	}
	t := data.Search("type").Data()
	if t == nil {
		log.Errorf("[ CATCHER ] Malformatted packet: %v", data.String())
		return
	}
	if msg.PktClass == "TELEMETRY" {
		StoragePipe <- data
	}
}

func StorageCatchingDaemon() {
	signal.Reserve(1)
	log.Debug("Entering catcher loop")
	for {
		if signal.ExitRequested() {
			break
		}
		if ! DoContinue {
			log.Debug("Programmatic loop exit")
			break
		}
		for len(StoragePipe) > 0 {
			log.Debugf("%v", len(StoragePipe))
			pkt := <- StoragePipe
	    if pkt == nil {
	      continue
	    }
			row := make([]tstorage.Row, 0)
			r := new(tstorage.Row)
			switch pkt.Search("type").Data().(string) {
			case "metric":
				val := float64(0)
				switch v := pkt.Search("metric", "value").Data().(type) {
				case float64:
					val = v
				case int64:
					val = float64(v)
				}
				labels := []tstorage.Label{
					{Name: "host", Value: pkt.Search("metric", "host").Data().(string)},
				}
				r.Metric = pkt.Search("metric", "key").Data().(string)
				r.Labels = labels
				r.DataPoint = tstorage.DataPoint{
					Timestamp: int64(pkt.Search("metric", "timestamp").Data().(float64)),
					Value:     val,
				}
			}
			row = append(row, *r)
			fmt.Println(row)
			err := Storage.InsertRows(row)
			if err != nil {
				log.Errorf("[ CATCHER ] Error storing telemetry: %v", err)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	signal.Release(1)
	log.Debug("Exiting catcher loop")
}

func InitStoragePipe() {
	log.Debug("Configuring internal telemetry storage pipeline")
	StoragePipe  = make(chan *gabs.Container, MAXPIPECAP)
	go StorageCatchingDaemon()
	NatsTelemetryRecv(StoreTelemetry)
}
