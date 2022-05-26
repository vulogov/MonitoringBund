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
	log.Debug("[ MBUND ] bund.Run() is reached")
	core := stdlib.InitBUND()
	for _, f := range *conf.Scripts {
		RunFile(core, f)
	}
	signal.ExitRequest()
}
