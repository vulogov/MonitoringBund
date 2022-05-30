package bund

import (
	"os"
	"context"
	"fmt"
	"strings"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/signal"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	"go.etcd.io/etcd/client/v3"
	"github.com/lrita/cmap"
	tc "github.com/vulogov/ThreadComputation"
)

var Etcd 							*clientv3.Client
var ApplicationId 		string
var ApplicationType 	string
var Conf  						cmap.Cmap

func ResetLocalConfig() {
	log.Debug("Reset local config cache")
	Conf.Range(func (key, value interface{}) bool {
		Conf.Delete(key)
		return true
	})
	SetDefaultConfiguration()
}

func SetDefaultConfiguration() {
	tc.SetVariable("Name", *conf.Name)
	tc.SetVariable("Id", *conf.Id)
	tc.SetVariable("ApplicationId", ApplicationId)
	tc.SetVariable("ApplicationType", ApplicationType)
}

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
	ApplicationType = otype
	SetApplicationId(otype)
}

func EtcdSetItem(key string, value string) {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	_, err := Etcd.Put(ctx, fmt.Sprintf("MBUND/%s/%s", *conf.Name, key), value)
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
}

func EtcdSetConfItem(key string, value string) {
	Conf.Store(key, value)
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	_key := fmt.Sprintf("MBUND/%s/conf/%s", *conf.Name, key)
	log.Debugf("EtcdSetConfItem(%v, %v)", _key, value)
	_, err := Etcd.Put(ctx, _key, value)
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
}

func EtcdGetConfItems() {
	ResetLocalConfig()
	log.Debug("Updating Local config cache from ETCD")
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	value, err := Etcd.Get(ctx, fmt.Sprintf("MBUND/%s/conf/", *conf.Name), clientv3.WithPrefix())
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	for _, v := range(value.Kvs) {
		key := strings.Split(string(v.Key), "/")
		log.Debugf("[ CONF ] %v = %v", key[len(key)-1], string(v.Value))
		Conf.Store(key[len(key)-1], v.Value)
	}
}

func EtcdReturnConfItems() *map[string]string {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	value, err := Etcd.Get(ctx, fmt.Sprintf("MBUND/%s/conf/", *conf.Name), clientv3.WithPrefix())
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

func EtcdGetItems()  *map[string]string {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	value, err := Etcd.Get(ctx, fmt.Sprintf("MBUND/%s/", *conf.Name), clientv3.WithPrefix())
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
	res := make(map[string]string)
	for _, v := range(value.Kvs) {
		key := strings.Split(string(v.Key), "/")
		if len(key) > 3 {
			continue
		}
		res[key[len(key)-1]] = string(v.Value)
	}
	return &res
}

func EtcdDelItems()   {
	ctx, _ := context.WithTimeout(context.Background(), *conf.Timeout)
	_, err := Etcd.Delete(ctx, fmt.Sprintf("MBUND/%s/", *conf.Name), clientv3.WithPrefix())
	if err != nil {
		log.Errorf("[ ETCD ] %v", err)
		signal.ExitRequest()
		os.Exit(10)
	}
}

func UpdateBundVariablesFromLocalConf(core *stdlib.BUNDEnv) {
	log.Debugf("Updating BUND environment with Local config")
	Conf.Range(func (key, value interface{}) bool {
		switch key.(type) {
		case string:
			core.TC.SetVariable(key.(string), value)
		}

		return true
	})
}

func UpdateLocalConfigFromEtcd() {
	ResetLocalConfig()
	etcd_cfg := EtcdGetItems()
	log.Debug("Updating local configuration from ETCD")
	*conf.Id = string((*etcd_cfg)["ID"])
	SetApplicationId(ApplicationType)
	log.Debugf("Application ID is %v", *conf.ApplicationId)
	if ! *conf.CNatsLocal {
		*conf.Gnats = (*etcd_cfg)["gnats"]
	}
	log.Debugf("NATS is %v", *conf.Gnats)
	EtcdGetConfItems()
	if ! *conf.CNatsLocal {
		if gnats, ok := Conf.Load("NATS"); ok {
			log.Debugf("Set NATS server address from CONF/NATS %v", string(gnats.([]byte)))
			*conf.Gnats = string(gnats.([]byte))
		}
	}
}

func UpdateConfigToEtcd() {
	if len(*conf.Etcd) > 0 {
		log.Debugf("Upload NRBUND configuration to ETCD")
		addr := (*conf.Etcd)[0]
		log.Debugf("Update ETCD endpoints with %s", addr)
		EtcdSetItem("etcd", addr)
		log.Debugf("Update GNATS endpoints with %s", *conf.Gnats)
		if ! *conf.CNatsLocal {
			EtcdSetItem("gnats", *conf.Gnats)
		}
		if *conf.CIdUpdated {
			EtcdSetItem("ID", *conf.Id)
		}
		log.Debugf("Updating CONF cache")
		Conf.Range(func (key, value interface{}) bool {
			switch key.(type) {
			case string:
				switch value.(type) {
				case string:
					log.Debugf("[ CONF ] Conf/%v=%v", key, value)
					EtcdSetConfItem(key.(string), value.(string))
				}
			}
			return true
		})
	}
}

func SetApplicationId(atype string) {
	if len(*conf.ApplicationId) == 0 {
		ApplicationId = fmt.Sprintf("%s:%s:%s", *conf.Id, *conf.Name, atype)
		*conf.ApplicationId = ApplicationId
	}
	log.Debugf("[ CONFIG ] SetApplicationId(%v)", ApplicationId)
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
