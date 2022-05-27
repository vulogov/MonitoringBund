package bund

import (
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/vulogov/monitoringbund/internal/stdlib"
)

func RunFile(core *stdlib.BUNDEnv, name string) {
	log.Debugf("Running: %v", name)
	code, err := tc.ReadFile(name)
	if err != nil {
		log.Fatalf("Error loading file: %v", err)
	}
	core.Eval(code)
}

func Run() {
	Init()
	InitEtcdAgent("run")
	UpdateLocalConfigFromEtcd()
	InitNatsAgent()
	if ! WaitSync() {
		return
	}
	log.Debugf("[ MBUND ] bund.Run(%v) is reached", ApplicationId)
	core := stdlib.InitBUND()
	for _, f := range *conf.Scripts {
		RunFile(core, f)
	}
	signal.ExitRequest()
}
