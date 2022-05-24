package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/davecgh/go-spew/spew"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/stdlib"
)

func NRBundSendToNewRelic(core *stdlib.BUNDEnv, desc map[string]interface{})  {
	if attrs, ok := desc["event"]; ok {
		NRBundSendCustomEvent(core, attrs.([]interface{}))
		return
	}
	if attrs, ok := desc["metric"]; ok {
		NRBundSendCustomMetric(core, attrs.([]interface{}))
		return
	}
}

func NRBundSendCustomEvent(core *stdlib.BUNDEnv, attrs []interface{}) {
	var data map[string]interface{}

	data = make(map[string]interface{})
	data["hostname"] = ApplicationId
	data["application"] = *conf.Name
	for _, k := range(attrs) {
		if core.TC.Ready() {
			data[k.(string)] = core.TC.Get()
		}
	}
	log.Debugf("EVT[%v] %v", *conf.EvtDst, spew.Sdump(data))
	NRapp.RecordCustomEvent(*conf.EvtDst, data)
}

func NRBundSendCustomMetric(core *stdlib.BUNDEnv, attrs []interface{}) {
	for _, k := range(attrs) {
		if core.TC.Ready() {
			val := core.TC.Get()
			switch val.(type) {
			case float64:
				log.Debugf("Sending float metric: %v = %v[%T]", k.(string), val, val)
				NRapp.RecordCustomMetric(k.(string), val.(float64))
			case int64:
				log.Debugf("Sending int metric: %v = %v[%T]", k.(string), val, val)
				NRapp.RecordCustomMetric(k.(string), float64(val.(int64)))
			default:
				log.Debugf("Not sending metric as data type is not supported: %v = %v[%T]", k.(string), val, val)
			}
		}
	}
}
