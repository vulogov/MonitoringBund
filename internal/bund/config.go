package bund

import (
	"fmt"
	"github.com/pieterclaerhout/go-log"
)



func Config() {
	Init()
	log.Debug("[ MBUND ] bund.Config() is reached")
	InitEtcdAgent("config")
	UpdateConfigToEtcd()
	fmt.Println(EtcdGetItems())
}
