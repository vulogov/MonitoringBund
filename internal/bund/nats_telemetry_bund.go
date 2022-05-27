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

func init() {
	stdlib.RegisterContextFunctionCallback("natssendmetric", tc.Json, bundNatsSendMetric)
	tc.SetFunction("metric.Send", BUNDNatsSendMetric)
}
