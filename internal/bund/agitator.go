package bund

import (
	"fmt"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/bamzi/jobrunner"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/pieterclaerhout/go-log"
)

type TheScript struct {
	Name    string
	Uri     string
	Res     map[string]interface{}
	Args    []string
}

func (s TheScript) Run() {
	if ! HadSync {
		log.Warn("Request to submit job received but agitator not in SYNC state. Request ignored.")
		return
	}
	txn := NRapp.StartTransaction(fmt.Sprintf("[%v]%v", s.Name, s.Uri))
	defer txn.End()
	log.Debugf("[ MBUND ] sending %s", s.Name)
	segment := txn.StartSegment(fmt.Sprintf("[%v] Read %v", s.Name, s.Uri))
	script, err := tc.ReadFile(s.Uri)
	segment.End()
	if err != nil {
		log.Errorf("[ MBUND ] %v", err)
		txn.NoticeError(err)
		return
	}
	if len(script) == 0 {
		log.Errorf("[ MBUND ] script can not be a zero length")
		txn.NoticeError(err)
		return
	}
	pkt, err := MakeScript(s.Uri, "agitator", []byte(script), s.Args, s.Res)
	NatsSend(pkt)
}


func AgitatorScheduleConfig() {
	txn := NRapp.StartTransaction(fmt.Sprintf("%s loading schedule configuration", ApplicationId))
	defer txn.End()
	for _, n := range(*conf.AConf) {
		segment := txn.StartSegment(fmt.Sprintf("%s loading %s", ApplicationId, n))
		cfg := HJsonLoadConfig(n)
		if cfg != nil {
			if jobs, ok := (*cfg)["jobs"]; ok {
				for _, j := range(jobs.([]interface{})) {
					job := j.(map[string]interface{})
					if name, ok := job["name"]; ok {
						if schedule, ok := job["schedule"]; ok {
							if uri, ok := job["uri"]; ok {
								sjob := TheScript{Name: name.(string), Uri: uri.(string)}
								if jargs, ok := job["args"]; ok {
									for _, _v := range(jargs.([]interface{})) {
										sjob.Args = append(sjob.Args, fmt.Sprintf("%v", _v))
									}
								}
								if sres, ok := job["return"]; ok {
									sjob.Res = sres.(map[string]interface{})
								}
								log.Debugf("Scheduling (%v)[%v]=%v", name.(string), schedule.(string), uri.(string))
								jobrunner.Schedule(schedule.(string), sjob)
							}
						}
					}
				}
			}
		}
		segment.End()
	}
}

func Agitator() {
	Init()
	log.Debug("[ MBUND ] bund.Agitator() is reached")
	InitEtcdAgent("agitator")
	if *conf.UploadConf {
		log.Info("Updating ETCD from local Agitator configuration")
		UpdateConfigToEtcd()
	} else {
		log.Info("Updating local Agitator configuration from ETCD")
		UpdateLocalConfigFromEtcd()
	}
	InitNatsAgent()
	InitNewRelicAgent()
	AgitatorScheduleConfig()
	jobrunner.Schedule("@every 5s", NATSSync{})
	Loop()
}
