package bund

import (
	"fmt"
	"os"
	"bufio"
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	"github.com/mgutz/ansi"
)

func EvalDisplayResult(core *stdlib.BUNDEnv) {
	var out string
	if core.TC.Ready() {
		e := core.TC.Get()
		core.TC.Set(e)
		fun := tc.GetConverterCallback(e)
		if fun == nil {
			out = fmt.Sprintf("%v", e)
		} else {
			out_add := fun(e, tc.String)
			if out_add == nil {
				out += fmt.Sprintf("%v", e)
			} else {
				out += out_add.(string)
			}
		}
		if *conf.ShowEResult {
			if *conf.Color {
				out = ansi.Color(out, "yellow")
				fmt.Println(out)
			} else {
				fmt.Println(out)
			}
		} else {
			log.Debugf("Result: %v", out)
		}
	} else {
		log.Debug("Stack is too shallow for result display")
	}
}

func BundEvalExpression(code string) {
	if *conf.CDebug {
		log.Info("BUND core debug is on")
		tc.SetVariable("tc.Debuglevel", "debug")
		log.Infof("[ MBUND ] core version: %v", tc.VERSION)
	} else {
		log.Debug("BUND core debug is off")
		tc.SetVariable("tc.Debuglevel", "info")
		log.Debugf("[ MBUND ] core version: %v", tc.VERSION)
	}
	core := stdlib.InitBUND()
	core.Eval(code)
	EvalDisplayResult(core)
}

func Eval() {
	Init()
	log.Debug("[ MBUND ] bund.Eval() is reached")
	if len(*conf.Expr) > 0 {
		log.Debugf("Evaluating expression from command line: %v", *conf.Expr)
		BundEvalExpression(*conf.Expr)
	} else if *conf.EStdin {
		code := ""
		log.Debug("Evaluating expression from STDIN")
		r := bufio.NewScanner(os.Stdin)
		for r.Scan() {
			code += r.Text()
			code += "\n"
		}
		if err := r.Err(); err != nil {
      log.Errorf("Error reading from STDIN: %v", err)
			return
    }
		BundEvalExpression(code)
	} else {
		log.Error("Evaluation expression not defined")
	}
	signal.ExitRequest()
}
