package bund

import (
	"fmt"
	"strings"
	"github.com/lrita/cmap"
	"github.com/peterh/liner"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/banner"
	"github.com/vulogov/monitoringbund/internal/stdlib"
)

var (
	shellCmd cmap.Cmap
	commands = []string{
		".version", ".exit", ".stack", ".last",
	}
	PROMPT = "[ MBUND ] "
)

func Shell() {
	Init()
	banner.PrintBanner(fmt.Sprintf("[ MBUND %v ]", conf.EVersion))
	log.Info("For exit, type: .exit")
	log.Debug("[ MBUND ] bund.Shell() is reached")
	line := liner.NewLiner()
	defer line.Close()
	line.SetCtrlCAborts(true)

	line.SetCompleter(func(line string) (c []string) {
		for _, n := range commands {
			if strings.HasPrefix(n, strings.ToLower(line)) {
				c = append(c, n)
			}
		}
		return
	})

	core := stdlib.InitBUND()

	out:
	for {
		if cmd, err := line.Prompt(PROMPT); err == nil {
			cmd = strings.Trim(cmd, "\n \t\r")
			line.AppendHistory(cmd)
			log.Debugf("shell get: %v", cmd)
			switch cmd {
			case ".exit":
				log.Debug("Exiting")
				break out
			default:
				if IsShellCommand(cmd) {
					log.Debugf("Running shell command: %v", cmd)
					RunShellCommand(cmd, core.TC)
				} else {
					log.Debug("Executing in ThreadComputation")
					core.Eval(cmd)
					ShellDisplayResult(core.TC, false)
					if core.TC.ExitRequested() {
						log.Debug("Exiting from shell")
						break out
					}
				}
			}
		} else if err == liner.ErrPromptAborted {
			log.Debug("Aborted")
			break
		} else {
			log.Debugf("Error reading line: %v", err)
		}
	}
}
