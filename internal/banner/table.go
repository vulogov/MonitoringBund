package banner

import (
	"os"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/tomlazar/table"
	tc "github.com/vulogov/ThreadComputation"

	"github.com/vulogov/monitoringbund/internal/conf"
)

func Table(display bool) {
	var cfg table.Config

	if !*conf.VTable && ! display {
		return
	}

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
	if *conf.VTable || display {
		tab := table.Table{
			Headers: []string{"Description", "Value"},
			Rows: [][]string{
				{"Name", *conf.Name},
				{"Version", conf.EVersion},
				{"Extended version", conf.BVersion},
				{"Core version", tc.VERSION},
				{"Args", *conf.Args},
				{"Core debug mode", fmt.Sprintf("%v", *conf.CDebug)},
			},
		}
		tab.WriteTable(os.Stdout, &cfg)
	}
}
