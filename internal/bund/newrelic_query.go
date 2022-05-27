package bund

import (
	"time"
	"errors"
	"github.com/google/uuid"
	"github.com/vulogov/nrapi"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/stdlib"
  "github.com/gammazero/deque"
  tc "github.com/vulogov/ThreadComputation"
)

func bundNewrelicQuery(l *tc.TCExecListener, v interface{}) interface{} {
	switch query := v.(type) {
	case string:
		log.Debugf("[ QUERY ] %v", query)
		_res := NRAPI.NRQL(query)
		if _res == nil {
			l.TC.SetError("[ QUERY ] \"%v\" had failed", query)
			return nil
		}
		df, err := nrapi.DataFrame(_res)
		if err != nil {
			l.TC.SetError("[ QUERY ] dataframe conversion failed: %v", err)
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

func BUNDNewrelicQuery(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	if ! NewRelicConfigured {
		log.Debug("New Relic connector not configured. Trying to configure")
		if ! InitNewRelicAgent() {
			return nil, errors.New("New Relic connector not configured")
		}
	}
	err := stdlib.ExecuteSingleArgumentFunction(l, "newrelicquery", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}


func init() {
	stdlib.RegisterContextFunctionCallback("newrelicquery", tc.String, bundNewrelicQuery)
	tc.SetFunction("nr.Query", BUNDNewrelicQuery)
}
