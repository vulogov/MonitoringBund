package bund

import (
	"bufio"
	"os"
	"github.com/vulogov/monitoringbund/internal/conf"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/pieterclaerhout/go-log"
	"github.com/hjson/hjson-go"
)



func Submit() {
	var err error

	Init()
	InitEtcdAgent("submit")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	log.Debugf("[ MBUND ] bund.Submit(%v) is reached", ApplicationId)
	script := ""
	if *conf.SScript == "--" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
        script += scanner.Text()
				script += "\n"
    }
	} else {
		script, err = tc.ReadFile(*conf.SScript)
		if err != nil {
			log.Errorf("[ MBUND ] %v", err)
			return
		}
	}
	if len(script) == 0 {
		log.Errorf("[ MBUND ] script can not be a zero length")
		return
	}
	res := new(map[string]interface{})
	if len(*conf.SReturn) > 0 {
		if err = hjson.Unmarshal([]byte(*conf.SReturn), res); err != nil {
			log.Errorf("[ CONFIG ] %v", err)
			return
	  }
	}
	pkt, err := MakeScript(*conf.SScript, "submit", []byte(script), *conf.SArgs, *res)
	if err != nil {
		log.Errorf("[ MBUND ] %v", err)
		return
	}
	log.Debugf("[ MBUND ] Sending script for execution len()=%v", len(pkt))
	NatsSend(pkt)
}
