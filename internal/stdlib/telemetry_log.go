package stdlib

import (
  "time"
  "errors"
  "github.com/Jeffail/gabs/v2"
  "github.com/gammazero/deque"
  "github.com/vulogov/monitoringbund/internal/conf"
  tc "github.com/vulogov/ThreadComputation"
)


func BUNDTelemetryLog(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var err error
  stamp := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
  out := gabs.New()
  out.Set("log", "type")
  out.Set(stamp, "log", "timestamp")
  out.Set(*conf.ApplicationId, "attributes", "ApplicationId")
  out.Set(*conf.Name, "attributes", "ApplicationName")
  for q.Len() > 0 {
    v := q.PopFront()
    switch data := v.(type) {
    case *tc.TCPair:
      switch key := data.X.(type) {
      case string:
        switch key {
        case "service", "host", "logtype":
          out.Set(data.Y, "log", key)
        default:
          out.Set(data.Y, "attributes", key)
        }
      }
    case string:
      out.Set(data, "log", "msg")
    default:
      out.Set(data, "attributes", "value")
    }
  }
  if out.Search("log", "service").Data() == nil {
    svc := l.TC.GetContext("key")
    if svc == nil {
      svc = "genericservice"
    }
    switch svc.(type) {
    case string:
      out.Set(svc, "log", "service")
    }
  }
  if out.Search("log", "logtype").Data() == nil {
    lt := l.TC.GetContext("logtype")
    if lt == nil {
      lt = "log"
    }
    switch lt.(type) {
    case string:
      out.Set(lt, "log", "logtype")
    }
  }
  if out.Search("log", "host").Data() == nil {
    host := l.TC.GetContext("host")
    if host == nil {
      host, err = tc.GetVariable("system.Hostname")
      if err != nil {
        return nil, err
      }
    }
    switch host.(type) {
    case string:
      out.Set(host, "log", "host")
    default:
      return nil, errors.New("host attribute for Event is not a string")
    }
  }
  res := new(tc.TCJson)
  res.J = out
  return res, nil
}

func init() {
  tc.SetFunction("Log", BUNDTelemetryLog)
}
