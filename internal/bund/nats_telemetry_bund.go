package bund

import (
	"fmt"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	"github.com/gammazero/deque"
	tc "github.com/vulogov/ThreadComputation"
)

func bundNatsSendMetric(l *tc.TCExecListener, v interface{}) interface{} {
	switch data := v.(type) {
	case *tc.TCJson:
		err := NatsSendMetric(data)
		if err != nil {
			l.TC.SetError(fmt.Sprintf("[ METRIC ] %v", err))
		}
	}
	return nil
}

func BUNDNatsSendMetric(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := stdlib.ExecuteSingleArgumentFunction(l, "natssendmetric", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func bundNatsSendEvent(l *tc.TCExecListener, v interface{}) interface{} {
	switch data := v.(type) {
	case *tc.TCJson:
		err := NatsSendEvent(data)
		if err != nil {
			l.TC.SetError(fmt.Sprintf("[ EVENT ] %v", err))
		}
	}
	return nil
}

func BUNDNatsSendEvent(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := stdlib.ExecuteSingleArgumentFunction(l, "natssendevent", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	stdlib.RegisterContextFunctionCallback("natssendmetric", tc.Json, bundNatsSendMetric)
	tc.SetFunction("metric.Send", BUNDNatsSendMetric)
	stdlib.RegisterContextFunctionCallback("natssendevent", tc.Json, bundNatsSendEvent)
	tc.SetFunction("event.Send", BUNDNatsSendEvent)

}
