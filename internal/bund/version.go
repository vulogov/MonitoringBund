package bund

import (
	"fmt"

	"github.com/pieterclaerhout/go-log"

	"github.com/vulogov/monitoringbund/internal/banner"
	"github.com/vulogov/monitoringbund/internal/conf"
)

func Version() {
	Init()
	log.Debug("[ MBUND ] bund.Version() is reached")
	banner.Banner(fmt.Sprintf("[ MBUND %v ]", conf.EVersion))
	banner.Table(true)
}
