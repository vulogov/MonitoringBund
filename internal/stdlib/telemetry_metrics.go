package stdlib

import (
  "time"
  "errors"
  "github.com/Jeffail/gabs/v2"
  "github.com/gammazero/deque"
  "github.com/vulogov/monitoringbund/internal/conf"
  tc "github.com/vulogov/ThreadComputation"
)


func BUNDTelemetryMetrics(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var err error
  stamp := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
  out := gabs.New()
  out.Set("metric", "type")
  out.Set(stamp, "metric", "timestamp")
  out.Set(*conf.ApplicationId, "attributes", "ApplicationId")
  out.Set(*conf.Name, "attributes", "ApplicationName")
  for q.Len() > 0 {
    v := q.PopFront()
    switch data := v.(type) {
    case *tc.TCPair:
      switch key := data.X.(type) {
      case string:
        switch key {
        case "key", "host":
          out.Set(data.Y, "metric", key)
        default:
          out.Set(data.Y, "attributes", key)
        }
      }
    default:
      out.Set(data, "metric", "value")
    }
  }
  if out.Search("metric", "key").Data() == nil {
    key := l.TC.GetContext("key")
    if key == nil {
      return nil, errors.New("Telemetry key not defined")
    }
    switch key.(type) {
    case string:
      out.Set(key, "metric", "key")
    }
  }
  if out.Search("metric", "host").Data() == nil {
    host := l.TC.GetContext("host")
    if host == nil {
      host, err = tc.GetVariable("system.Hostname")
      if err != nil {
        return nil, err
      }
    }
    switch host.(type) {
    case string:
      out.Set(host, "metric", "host")
    default:
      return nil, errors.New("host attribute for Metric is not a string")
    }
  }
  if out.Search("metric", "type").Data() == nil {
    mtype := l.TC.GetContext("type")
    if mtype == nil {
      mtype = "gauge"
    }
    switch mtype.(type) {
    case string:
      out.Set(mtype, "metric", "type")
    default:
      return nil, errors.New("Metric Type attribute for Metric is not a string")
    }
  }
  res := new(tc.TCJson)
  res.J = out
  return res, nil
}

func init() {
  tc.SetFunction("Metric", BUNDTelemetryMetrics)
}
