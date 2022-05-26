package bund

import (
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/stdlib"
)


func BundGlobalEvalExpression(code string, args []string, res map[string]interface{}) {
	if *conf.CDebug {
		log.Info("BUND core debug is on")
		tc.SetVariable("tc.Debuglevel", "debug")
		log.Infof("[ MBUND ] core version: %v", tc.VERSION)
	} else {
		log.Debug("BUND core debug is off")
		tc.SetVariable("tc.Debuglevel", "info")
		log.Debugf("[ MBUND ] core version: %v", tc.VERSION)
	}
	log.Debugf("BUND core display result %v", *conf.ShowResult)
	core := stdlib.InitBUND()
	if args != nil && len(args) > 0 {
		log.Debugf("[ MBUND ] Args=len(%v)", len(args))
	}
	if len(args) > 0 {
		core.TC.EvAttrs.PushFront(AttrsToQueue(args))
	}
	UpdateBundVariablesFromLocalConf(core)
	core.Eval(code)
	if len(args) > 0 {
		core.TC.EvAttrs.PopFront()
	}
}
