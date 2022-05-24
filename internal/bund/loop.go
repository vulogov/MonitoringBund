package bund

import (
	"time"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/signal"
)

var DoContinue bool

func Loop() {
	log.Debug("Entering event loop")
	for {
		if signal.ExitRequested() {
			break
		}
		if ! DoContinue {
			log.Debug("Programmatic exent loop exit")
			break
		}
		time.Sleep(1 * time.Second)
	}
	log.Debug("Exiting event loop")
}

func init() {
	DoContinue = true
}
