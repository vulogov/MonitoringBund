package bund

import (

	"github.com/nakabonne/tstorage"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/vulogov/monitoringbund/internal/conf"
)

var Storage tstorage.Storage
var StorageConfigured bool

func InitInternalStorage() {
	var err error

	log.Debugf("Initializing internal metric storage")
	Storage, err = tstorage.NewStorage(
		tstorage.WithWriteTimeout(*conf.Timeout),
		tstorage.WithRetention(*conf.Retention),
	)
	if err != nil {
		log.Errorf("Internal storage error: %v", err)
		signal.ExitRequest()
	}
	StorageConfigured = true
}

func CloseInternalStorage() {
	log.Debugf("Closing internal metric storage")
	if StorageConfigured {
		Storage.Close()
	}
}

func init() {
	StorageConfigured = false
}
