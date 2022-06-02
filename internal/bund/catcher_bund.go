package bund

import (
	"fmt"
	"time"
	"github.com/gammazero/deque"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	"github.com/nakabonne/tstorage"
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
)

func bundQueryStorage(l *tc.TCExecListener, v interface{}) interface{} {
	switch data := v.(type) {
	case *tc.TCPair:
		switch host := data.X.(type) {
		case string:
			switch key := data.Y.(type) {
			case string:
				stamp1 := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
				stamp2 := stamp1 - (int64((*conf.Retention).Seconds())*1000)
				labels := []tstorage.Label{
					{Name: "host", Value: host},
				}
				log.Debugf("metric.Get(%v %v %v %v)", key, labels, stamp2, stamp1)
				res, err := Storage.Select(key, labels, stamp2, stamp1)
				if err != nil {
					l.TC.SetError(fmt.Sprintf("[ STORAGE ]: %v", err))
					return nil
				}
				out := tc.MakeList()
				for _, v := range res {
					e := l.TC.Pair(v.Timestamp, v.Value)
					out.Q.PushBack(e)
				}
				return out
			}
		}
	}
	return nil
}

func BUNDQueryStorage(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := stdlib.ExecuteSingleArgumentFunction(l, "storage", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func bundSampleStorage(l *tc.TCExecListener, v interface{}) interface{} {
	switch data := v.(type) {
	case *tc.TCPair:
		switch host := data.X.(type) {
		case string:
			switch key := data.Y.(type) {
			case string:
				stamp1 := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
				stamp2 := stamp1 - (int64((*conf.Retention).Seconds())*1000)
				labels := []tstorage.Label{
					{Name: "host", Value: host},
				}
				res, err := Storage.Select(key, labels, stamp2, stamp1)
				if err != nil {
					l.TC.SetError(fmt.Sprintf("[ STORAGE ]: %v", err))
					return nil
				}
				out := tc.MakeNumbers()
				for _, v := range res {
					out.Add(v.Value)
				}
				return out
			}
		}
	}
	return nil
}

func BUNDSampleStorage(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := stdlib.ExecuteSingleArgumentFunction(l, "storagesample", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	stdlib.RegisterContextFunctionCallback("storage", tc.Pair, bundQueryStorage)
	tc.SetFunction("metric.Get", BUNDQueryStorage)
	stdlib.RegisterContextFunctionCallback("storagesample", tc.Pair, bundSampleStorage)
	tc.SetFunction("metric.Sample", BUNDSampleStorage)
}
