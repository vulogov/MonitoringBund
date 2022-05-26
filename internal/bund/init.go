package bund

import (
	"github.com/cosiner/argv"
	"github.com/pieterclaerhout/go-log"
	"github.com/bamzi/jobrunner"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	tlog "github.com/vulogov/monitoringbund/internal/log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/monitoringbund/internal/signal"
)

func Init() {
	tlog.Init()
	stdlib.InitStdlib()
	log.Debug("[ MBUND ] bund.Init() is reached")
	signal.InitSignal()
	if len(*conf.Args) > 0 {
		Argv, err := argv.Argv(*conf.Args, func(backquoted string) (string, error) {
			return backquoted, nil
		}, nil)
		if err != nil {
			log.Fatalf("Error parsing ARGS: %v", err)
		}
		log.Debugf("ARGV: %v", Argv)
		conf.Argv = Argv
	}
	log.Debugf("[ MBUND ] Id: %v", *conf.Id)
	log.Debugf("[ MBUND ] Name: %v", *conf.Name)
	jobrunner.Start(*conf.JPool, *conf.JCon)
	log.Debugf("[ MBUND ] Job runner started")
	stdlib.StoreArgs()
	if *conf.CDebug {
		log.Info("BUND core debug is on")
		tc.SetVariable("tc.Debuglevel", "debug")
		log.Infof("[ MBUND ] core version: %v", tc.VERSION)
	} else {
		log.Debug("BUND core debug is off")
		tc.SetVariable("tc.Debuglevel", "info")
		log.Debugf("[ MBUND ] core version: %v", tc.VERSION)
	}
	if *conf.CNatsLocal {
		log.Info("NATS configuration will not be taken from ETCD")
	} else {
		log.Debug("ETCD will be probed for NATS configuration")
	}
}
