package bund

import (
	"time"
	"errors"
	"github.com/google/uuid"
	// "github.com/prometheus/client_golang/api"
	// v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	"github.com/gammazero/deque"
	tc "github.com/vulogov/ThreadComputation"
)

func bundPrometheusQuery(l *tc.TCExecListener, v interface{}) interface{} {
	switch query := v.(type) {
	case string:
		log.Debugf("[ QUERY ] %v", query)

		res := new(tc.TCData)
	  res.Id    = uuid.NewString()
	  res.Stamp = time.Now()
		// res.D     = df
		return res
	}
	return nil
}

func BUNDPrometheusQuery(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	if ! PrometheusConfigured {
		log.Debug("Prometheus connector not configured. Trying to configure")
		if ! InitPrometheusAgent() {
			return nil, errors.New("Prometheus connector not configured")
		}
	}
	err := stdlib.ExecuteSingleArgumentFunction(l, "prometheusquery", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	stdlib.RegisterContextFunctionCallback("prometheusquery", tc.String, bundPrometheusQuery)
	tc.SetFunction("prometheus.Query", BUNDPrometheusQuery)
}
