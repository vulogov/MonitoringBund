package bund

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/gammazero/deque"
	"github.com/vulogov/monitoringbund/internal/stdlib"
	tc "github.com/vulogov/ThreadComputation"
)

func bundSetLocalConfig(l *tc.TCExecListener, v interface{}) interface{} {
  switch src := v.(type) {
  case *tc.TCPair:
		switch key := src.X.(type) {
		case string:
			switch val := src.Y.(type) {
			case string:
				log.Debugf("[ CONF ] Conf.Store(%v, %v)", key, val)
				Conf.Store(key, val)
			}
		}
  }
  return nil
}

func BUNDSetLocalConf(l *tc.TCExecListener, name string, q *deque.Deque) (interface{}, error) {
	err := stdlib.ExecuteSingleArgumentFunction(l, "localconfigset", q)
  if err != nil {
    return nil, err
  }
	return nil, nil
}

func init() {
	stdlib.RegisterContextFunctionCallback("localconfigset", tc.Pair, bundSetLocalConfig)
	tc.SetCommand("local.SetConfig", BUNDSetLocalConf)
}
