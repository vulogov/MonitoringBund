package bund

import (
	"os"
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/tomlazar/table"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/stdlib"
)

func DisplayEtcdConfig() {
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
	m_range := EtcdGetItems()
	c_range := EtcdReturnConfItems()
	for k, v := range  *m_range {
		data = append(data, []string{k, v})
	}
	for k, v := range  *c_range {
		data = append(data, []string{fmt.Sprintf("CONF/%v", k), v})
	}
	tab := table.Table{
		Headers: []string{"Key", "Value"},
		Rows: data,
	}
	tab.WriteTable(os.Stdout, &cfg)
}

func Config() {
	Init()
	log.Debugf("[ MBUND ] bund.Config(%v) is reached", ApplicationId)
	InitEtcdAgent("config")
	core := stdlib.InitBUND()
	for _, n := range(*conf.SConf) {
		log.Debugf("[ CONF ] Processing %v", n)
		RunFile(core, n)
	}
	UpdateConfigToEtcd()
	if *conf.CShow {
		DisplayEtcdConfig()
	}
	UpdateLocalConfigFromEtcd()
}
