package bund

import (
	"time"
	"github.com/Jeffail/gabs/v2"
	"github.com/vulogov/monitoringbund/internal/conf"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/pieterclaerhout/go-log"
)



func Telemetry() {
	var err error

	Init()
	InitEtcdAgent("telemetry")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! WaitSync() {
		return
	}
	log.Debugf("[ MBUND ] bund.Telemetry(%v) is reached", ApplicationId)
	out := gabs.New()
	stamp := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
	out.Set(*conf.TType, "type")
	out.Set(*conf.ApplicationId, "attributes", "ApplicationId")
	out.Set(*conf.Name, "attributes", "ApplicationName")
	out.Set(stamp, *conf.TType, "timestamp")
	out.Set(*conf.TKey, *conf.TType, "key")
	out.Set(*conf.THost, *conf.TType, "host")
	switch *conf.TType {
	case "event":
		out.Set(tc.GetSimpleData(*conf.TValue), "attributes", "value")
		out.Set(*conf.TDst, "event", "destination")
	case "metric":
		out.Set(tc.GetSimpleData(*conf.TValue), "metric", "value")
		out.Set(*conf.TMType, "metric", "type")
	case "log":
		out.Set(*conf.TValue, "log", "msg")
		out.Set(*conf.TLSrv, "log", "service")
		out.Set(*conf.TLLt, "log", "logtype")
	default:
		log.Errorf("Unknown telemetry type: %v", *conf.TType)
		return
	}
	for k, v := range *conf.TArgs {
		out.Set(v, "attributes", k)
	}
	res := new(tc.TCJson)
  res.J = out
	switch *conf.TType {
	case "event":
		err = NatsSendEvent(res)
	case "log":
		err = NatsSendLog(res)
	case "metric":
		err = NatsSendMetric(res)
	}
	if err != nil {
		log.Errorf("[ TELEMETRY ]: %v", err)
	}
}
