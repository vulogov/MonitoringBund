package bund

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/monitoringbund/internal/conf"
)

func ShellDisplayResult(core *tc.TCstate, show bool) {
	var out string
	if core.Ready() {
		e := core.Get()
		core.Set(e)
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
		if *conf.ShowSResult || show {
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
