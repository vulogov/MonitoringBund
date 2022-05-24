package bund

import (
	"os"
	"context"
	"fmt"
	"strings"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/signal"
	"go.etcd.io/etcd/client/v3"
)

var Etcd *clientv3.Client
var ApplicationId string
var ApplicattionType string

func InitEtcdAgent(otype string) {
	var err error

	log.Debugf("Connecting to ETCD: %v", *conf.Etcd)
	Etcd, err = clientv3.New(
		clientv3.Config{
			Endpoints: *conf.Etcd,
			DialTimeout: *conf.Timeout,
		})
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	log.Debug("Sync ETCD endpoints")
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	err = Etcd.Sync(ctx)
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	ApplicattionType = otype
	SetApplicationId(otype)
}

func EtcdSetItem(key string, value string) {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	_, err := Etcd.Put(ctx, fmt.Sprintf("NRBUND/%s/%s", *conf.NRAccount, key), value)
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
}

func EtcdGetItems()  *map[string]string {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	value, err := Etcd.Get(ctx, fmt.Sprintf("NRBUND/%s/", *conf.NRAccount), clientv3.WithPrefix())
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	res := make(map[string]string)
	for _, v := range(value.Kvs) {
		key := strings.Split(string(v.Key), "/")
		res[key[len(key)-1]] = string(v.Value)
	}
	return &res
}

func UpdateLocalConfigFromEtcd() {
	etcd_cfg := EtcdGetItems()
	log.Debug("Updating local configuration from ETCD")
	*conf.Id = (*etcd_cfg)["ID"]
	SetApplicationId(ApplicattionType)
	log.Debugf("Application ID is %v", *conf.Id)
	log.Debugf("NR Account is %v", *conf.NRAccount)
	*conf.NRKey = (*etcd_cfg)["NEWRELIC_API_KEY"]
	*conf.NRLicenseKey = (*etcd_cfg)["NEWRELIC_LICENSE_KEY"]
	*conf.NRIngestKey = (*etcd_cfg)["NEWRELIC_INGEST_KEY"]
	*conf.Gnats = (*etcd_cfg)["gnats"]
	log.Debugf("NATS is %v", *conf.Gnats)
}

func UpdateConfigToEtcd() {
	if len(*conf.Etcd) > 0 {
		log.Debugf("Upload NRBUND configuration to ETCD")
		addr := (*conf.Etcd)[0]
		log.Debugf("Update ETCD endpoints with %s", addr)
		EtcdSetItem("etcd", addr)
		log.Debugf("Update GNATS endpoints with %s", *conf.Gnats)
		EtcdSetItem("gnats", *conf.Gnats)
		log.Debug("Upload NR keys")
		EtcdSetItem("NEWRELIC_ACCOUNT", *conf.NRAccount)
		EtcdSetItem("NEWRELIC_API_KEY", *conf.NRKey)
		EtcdSetItem("NEWRELIC_INGEST_KEY", *conf.NRIngestKey)
		EtcdSetItem("NEWRELIC_LICENSE_KEY", *conf.NRLicenseKey)
		EtcdSetItem("ID", *conf.Id)
	}
}

func SetApplicationId(atype string) {
	ApplicationId = fmt.Sprintf("%s:%s:%s", *conf.Id, *conf.Name, atype)
}

func CloseEtcdAgent() {
	log.Debug("Closing ETCD agent")
	if Etcd != nil {
		Etcd.Close()
	}
}

func init() {
	ApplicationId = "UNKNOWN"
}
