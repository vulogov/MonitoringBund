package bund

import (
	"fmt"
	"time"
	"strconv"
	"github.com/pieterclaerhout/go-log"
	"github.com/gammazero/deque"
	"github.com/Jeffail/gabs/v2"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	tc "github.com/vulogov/ThreadComputation"
)


func bundLoadStorage(l *tc.TCExecListener, v interface{}) interface{} {
	switch data := v.(type) {
	case *tc.TCData:
		host := l.TC.GetContext("host")
		key  := l.TC.GetContext("key")
		column := l.TC.GetContext("column")
		tscolumn := l.TC.GetContext("tscolumn")
		switch host.(type) {
		case string:
			switch key.(type) {
			case string:
				switch column.(type) {
				case string:
					switch tscolumn.(type) {
					case string:
			      for n := 0; n < data.D.NRows(); n++ {
							row := data.D.Row(n, false)
							if v, ok := row[column]; ok {
								if ts, ok := row[tscolumn]; ok {
									switch v.(type) {
									case float64:
										timestamp, err := strconv.ParseFloat(fmt.Sprintf("%v", ts), 64)
										if err != nil {
											log.Debugf("timestamp parsing error: %v", err)
											continue
										}

										pkt := gabs.New()
										pkt.Set("metric", "type")
										pkt.Set(key.(string), "metric", "key")
										pkt.Set(host.(string), "metric", "host")
										pkt.Set(timestamp, "metric", "timestamp")
										pkt.Set(v.(float64), "metric", "value")
										StoragePipe <- pkt
									}
								}
							}
			      }
						time.Sleep(5*time.Second)
					}
				}
			}
		}
	}
	return nil
}

func BUNDLoadStorage(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := stdlib.ExecuteSingleArgumentFunction(l, "storageload", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	stdlib.RegisterContextFunctionCallback("storageload", tc.Data, bundLoadStorage)
	tc.SetFunction("metric.Load", BUNDLoadStorage)
}
