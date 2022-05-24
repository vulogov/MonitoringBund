package bund

import (
	tc "github.com/vulogov/ThreadComputation"
	"github.com/pieterclaerhout/go-log"
	"github.com/hjson/hjson-go"
)

func HJsonLoadConfig(uri string) *map[string]interface{} {
	log.Debugf("[ CONFIG ] Loading %v", uri)
	cdat, err := tc.ReadFile(uri)
	if err != nil {
		log.Errorf("[ CONFIG ] %v", err)
		return nil
	}
	res := new(map[string]interface{})
	if err = hjson.Unmarshal([]byte(cdat), res); err != nil {
		log.Errorf("[ CONFIG ] %v", err)
		return nil
  }
	return res
}
