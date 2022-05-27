package stdlib

import (
  "time"
  "errors"
  "github.com/Jeffail/gabs/v2"
  "github.com/gammazero/deque"
  "github.com/vulogov/monitoringbund/internal/conf"
  tc "github.com/vulogov/ThreadComputation"
)


func BUNDTelemetryEvent(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
  var err error
  stamp := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
  out := gabs.New()
  out.Set("event", "type")
  out.Set(stamp, "event", "timestamp")
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
          out.Set(data.Y, "event", key)
        default:
          out.Set(data.Y, "attributes", key)
        }
      }
    default:
      out.Set(data, "attributes", "value")
    }
  }
  if out.Search("event", "key").Data() == nil {
    key := l.TC.GetContext("key")
    if key == nil {
      return nil, errors.New("Telemetry key not defined")
    }
    switch key.(type) {
    case string:
      out.Set(key, "event", "key")
    }
  }
  if out.Search("event", "host").Data() == nil {
    host := l.TC.GetContext("host")
    if host == nil {
      host, err = tc.GetVariable("system.Hostname")
      if err != nil {
        return nil, err
      }
    }
    switch host.(type) {
    case string:
      out.Set(host, "event", "host")
    default:
      return nil, errors.New("host attribute for Event is not a string")
    }
  }
  res := new(tc.TCJson)
  res.J = out
  return res, nil
}

func init() {
  tc.SetFunction("Event", BUNDTelemetryEvent)
}
