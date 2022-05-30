package bund

import (
	"os"
	"fmt"
	"time"
	"errors"
	"context"
	"github.com/google/uuid"
	"github.com/Jeffail/gabs/v2"
	"github.com/prometheus/common/model"
	"github.com/vulogov/nrapi"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	"github.com/gammazero/deque"
	tc "github.com/vulogov/ThreadComputation"
)

func bundPrometheusQuery(l *tc.TCExecListener, v interface{}) interface{} {
	switch query := v.(type) {
	case string:
		log.Debugf("[ QUERY ] %v", query)

		client, err := api.NewClient(api.Config{
			Address: *conf.PR_api_url,
		})
		if err != nil {
			fmt.Errorf("[ PROMETEUS ]: %v\n", err)
			os.Exit(1)
		}
		v1api := v1.NewAPI(client)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		// result, warnings, err := v1api.Query(ctx, query, time.Now(), v1.WithTimeout(*conf.Timeout))
		result, warnings, err := v1api.Query(ctx, query, time.Now())
		if err != nil {
			l.TC.SetError(fmt.Sprintf("[ PROMETHEUS ]: %v", err))
			return nil
		}
		if len(warnings) > 0 {
			log.Warnf("[ PROMETHEUS ]: %v", warnings)
		}
		container := gabs.New()
		container.Array()
		switch result.(type) {
    case model.Vector:
      vector := result.(model.Vector)
      for _, elem := range vector {
				je, _ := elem.MarshalJSON()
				src_row, _ := gabs.ParseJSON(je)
				row   := gabs.New()
				row.Set(src_row.S("value", "0"), "timestamp")
				row.Set(tc.GetSimpleData(src_row.S("value", "1").Data().(string)), "value")
				container.ArrayAppend(row)
			}
		case model.Matrix:
			matrix := result.(model.Matrix)
			for _, elem := range matrix {
				for _, e := range elem.Values {
					je, _ := e.MarshalJSON()
					src_row, _ := gabs.ParseJSON(je)
					row   := gabs.New()
					row.Set(src_row.S("0"), "timestamp")
					row.Set(tc.GetSimpleData(src_row.S("1").Data().(string)), "value")
					container.ArrayAppend(row)
				}
			}
		}
		df, err := nrapi.DataFrame(container)
		if err != nil {
			l.TC.SetError(fmt.Sprintf("[ PROMETHEUS ]: %v", err))
			return nil
		}
		res := new(tc.TCData)
	  res.Id    = uuid.NewString()
	  res.Stamp = time.Now()
		res.D     = df
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
