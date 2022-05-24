package bund

import (
	"fmt"
	"os"
	"github.com/mgutz/ansi"
	"github.com/tomlazar/table"
	"github.com/pieterclaerhout/go-log"
	tc "github.com/vulogov/ThreadComputation"
	"github.com/vulogov/monitoringbund/internal/banner"
	"github.com/vulogov/monitoringbund/internal/conf"
)

type ShellCommand func(*tc.TCstate) interface{}

func AddShellCommand(name string, fun ShellCommand) {
	shellCmd.Delete(name)
	shellCmd.Store(name, fun)
}

func RunShellCommand(name string, core *tc.TCstate) {
	if fun, ok := shellCmd.Load(name); ok {
		res := fun.(ShellCommand)(core)
		if res != nil {
			log.Infof("Returned: %v", res)
		}
	} else {
		log.Errorf("Shell command: %v not found", name)
	}
}

func IsShellCommand(name string) bool {
	if _, ok := shellCmd.Load(name); ok {
		return true
	}
	return false
}

func ShellCommandVersion(core *tc.TCstate) interface{} {
	banner.Table(true)
	return nil
}

func ShellCommandLast(core *tc.TCstate) interface{} {
	ShellDisplayResult(core, true)
	return nil
}

func ShellCommandStack(core *tc.TCstate) interface{} {
	var cfg table.Config
	var data [][]string

	cfg.ShowIndex = true
	if *conf.Color {
		cfg.Color = true
		cfg.AlternateColors = true
		cfg.TitleColorCode = ansi.ColorCode("white+buf")
		cfg.AltColorCodes = []string{"", ansi.ColorCode("white:grey+h")}
	} else {
		cfg.Color = false
		cfg.AlternateColors = false
		cfg.TitleColorCode = ansi.ColorCode("white+buf")
		cfg.AltColorCodes = []string{"", ansi.ColorCode("white:grey+h")}
	}
	if core.Ready() {
		for x := 0; x < core.Res.Len(); x++ {
			e := core.Res.Q().At(x)
			data = append(data, []string{fmt.Sprintf("%v", e)})
		}
	}
	tab := table.Table{
		Headers: []string{"Value"},
		Rows: data,
	}
	tab.WriteTable(os.Stdout, &cfg)
	return nil
}

func init() {
	AddShellCommand(".version", ShellCommandVersion)
	AddShellCommand(".stack", ShellCommandStack)
	AddShellCommand(".last", ShellCommandLast)
}
