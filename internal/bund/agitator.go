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
	log.Debugf("[ MBUND ] sending %s", s.Name)
	script, err := tc.ReadFile(s.Uri)
	if err != nil {
		log.Errorf("[ MBUND ] %v", err)
		return
	}
	if len(script) == 0 {
		log.Errorf("[ MBUND ] script can not be a zero length")
		return
	}
	pkt, err := MakeScript(s.Uri, "agitator", []byte(script), s.Args, s.Res)
	NatsSend(pkt)
}


func AgitatorScheduleConfig() {
	for _, n := range(*conf.AConf) {
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
	AgitatorScheduleConfig()
	Loop()
}
