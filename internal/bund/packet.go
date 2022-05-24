package bund

import (
	"github.com/google/uuid"
	"github.com/vmihailenco/msgpack"
	"github.com/pieterclaerhout/go-log"
	"github.com/vulogov/monitoringbund/internal/conf"
	"github.com/vulogov/monitoringbund/internal/signal"
)

type NRBundPacket struct {
	Id        string
	Uri       string
	PktId     string
	OrigName  string
	OrgRole   string
	PktClass  string
	PktKey    string
	Args      []string
	Res       map[string]interface{}
	Value     []byte
}

func NewBundPacket(uri string, orole string, pktclass string, pktkey string, args []string, value []byte, ret map[string]interface{}) *NRBundPacket {
	res := new(NRBundPacket)
	res.PktId = uuid.New().String()
	res.Id 				= *conf.Id
	res.Uri       = uri
	res.OrigName 	= *conf.Name
	res.OrgRole   = orole
	res.PktClass  = pktclass
	res.PktKey  	= pktkey
	res.Args      = args
	res.Res       = ret
	res.Value 		= value
	return res
}

func Marshal(uri string, orole string, pktclass string, pktkey string, args []string, value []byte, ret map[string]interface{}) ([]byte, error) {
	res := NewBundPacket(uri, orole, pktclass, pktkey, args, value, ret)
	return msgpack.Marshal(res)
}

func MarshalPacket(pkt *NRBundPacket) ([]byte, error) {
	return msgpack.Marshal(pkt)
}

func UnMarshal(data []byte) *NRBundPacket {
	res := new(NRBundPacket)
	err := msgpack.Unmarshal(data, res)
	if err != nil {
		log.Errorf("[ PACKET ] %v", err)
		return nil
	}
	return res
}

func IfSTOP(msg *NRBundPacket) bool {
	if msg.PktClass == "SYS" && msg.PktKey == "STOP" {
		signal.ExitRequest()
		DoContinue = false
		return true
	}
	return false
}

func IfSYNC(msg *NRBundPacket) bool {
	if msg.PktClass == "SYS" && msg.PktKey == "SYNC" {
		return true
	}
	return false
}

func MakeSync(orole string) ([]byte, error) {
	return Marshal("", orole, "SYS", "SYNC", nil, nil, nil)
}

func MakeStop(orole string) ([]byte, error) {
	return Marshal("", orole, "SYS", "STOP", nil, nil, nil)
}

func MakeScript(uri string, orole string, script []byte, args []string, res map[string]interface{}) ([]byte, error) {
	return Marshal(uri, orole, "SYS", "Agitator", args, script, res)
}
