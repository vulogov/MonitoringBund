package bund

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/vulogov/monitoringbund/internal/conf"
)

func Main() {
	switch kingpin.MustParse(conf.App.Parse(os.Args[1:])) {
	case conf.Version.FullCommand():
		Version()
	case conf.Shell.FullCommand():
		Shell()
	case conf.Run.FullCommand():
		Run()
	case conf.Eval.FullCommand():
		Eval()
	case conf.Agitator.FullCommand():
		Agitator()
	case conf.Agent.FullCommand():
		Agent()
	case conf.Config.FullCommand():
		Config()
	case conf.Submit.FullCommand():
		Submit()
	case conf.Take.FullCommand():
		Take()
	case conf.Sync.FullCommand():
		Sync()
	case conf.Watch.FullCommand():
		Watch()
	case conf.Stop.FullCommand():
		Stop()
	case conf.NRClient.FullCommand():
		Newrelic_Client()
	case conf.NRQLshell.FullCommand():
		Newrelic_NRQL_Shell()
	case conf.ZBXClient.FullCommand():
		Zabbix_Client()
	case conf.PRClient.FullCommand():
		Prometheus_Client()
	}
	Fin()
}
