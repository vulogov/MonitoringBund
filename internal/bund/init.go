package bund

import (
	"github.com/cosiner/argv"
	"github.com/pieterclaerhout/go-log"
	"github.com/bamzi/jobrunner"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	tlog "github.com/vulogov/monitoringbund/internal/log"
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
}
