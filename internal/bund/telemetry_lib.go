package bund

import (
	"fmt"
	"github.com/Jeffail/gabs/v2"
)

func TelemetryAttributesToMap(data *gabs.Container) map[string]interface{} {
	fmt.Println(data.String())
	res := make(map[string]interface{})
	for key, value := range data.S("attributes").ChildrenMap() {
		res[key] = value
	}
	return res
}
